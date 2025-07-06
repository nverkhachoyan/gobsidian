package generators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/charmbracelet/log"

	"path/filepath"
	"sort"
	"strings"

	"gobsidian/internal/executors"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/repository"
	"gobsidian/internal/terminal"
	"gobsidian/internal/transformers"
	"gobsidian/internal/utils"

	"github.com/gomarkdown/markdown/html"

	"github.com/gomarkdown/markdown"
)

type Generator interface {
	Generate() error
}

type StaticSiteGenerator struct {
	InputDirectory         string
	OutputDirectory        string
	SiteTitle              string
	SiteSubtitle           string
	BaseURL                string
	NotesPerPage           int
	Regexes                *models.ParserRegexes
	Logger                 *log.Logger
	Templates              *template.Template
	Parser                 parsers.Parser
	MarkdownRenderer       *html.Renderer
	shouldPrintVaultHealth bool
}

type PaginationData struct {
	HasPrev     bool
	PrevPageURL string
	HasNext     bool
	NextPageURL string
	CurrentPage int
	TotalPages  int
}

type BaseTemplateData struct {
	SiteTitle    string
	SiteSubtitle string
	BaseURL      string
	CurrentYear  int
	FileTree     *models.Folder
	Graph        template.JS
	Tags         []models.Tag
}

func NewStaticSiteGenerator(
	inputDirectory string,
	outputDirectory string,
	siteTitle string,
	siteSubtitle string,
	baseURL string,
	notesPerPage int,
	regexes *models.ParserRegexes,
	logger *log.Logger,
	templates *template.Template,
	parser parsers.Parser,
	markdownRenderer *html.Renderer,
	shouldPrintVaultHealth bool,
) (*StaticSiteGenerator, error) {
	return &StaticSiteGenerator{
		InputDirectory:         inputDirectory,
		OutputDirectory:        outputDirectory,
		SiteTitle:              siteTitle,
		SiteSubtitle:           siteSubtitle,
		BaseURL:                baseURL,
		NotesPerPage:           notesPerPage,
		Regexes:                regexes,
		Logger:                 logger,
		Templates:              templates,
		Parser:                 parser,
		MarkdownRenderer:       markdownRenderer,
		shouldPrintVaultHealth: shouldPrintVaultHealth,
	}, nil
}

func (g *StaticSiteGenerator) Generate() error {
	start := time.Now()

	stepDefs := []terminal.StepDefinition{
		terminal.CreateStep("Scanning files", "📄"),
		terminal.CreateStep("Parsing notes", "📝"),
		terminal.CreateStep("Building graph", "🔗"),
		terminal.CreateStep("Transforming content", "✨"),
		terminal.CreateStep("Building folder tree", "📁"),
		terminal.CreateStep("Executing templates", "🎨"),
		terminal.CreateStep("Copying assets", "📦"),
	}

	progressManager := terminal.NewProgressManager(stepDefs)
	progressManager.StartAsync()

	// Step 1: Scan notes
	stepStart := time.Now()
	noteScanner := NewNoteScanner(g.Logger, g.InputDirectory)
	scanTime, err := noteScanner.ScanAllNotes()
	if err != nil {
		progressManager.SendStepComplete(0, time.Since(stepStart), 0, err)
		return err
	}
	scannedNotes := noteScanner.GetAllNotes()
	progressManager.SendStepComplete(0, scanTime, len(scannedNotes), nil)

	// Step 2: Parse notes
	stepStart = time.Now()
	notesRepository := repository.NewNoteRepository()
	graphGenerator := NewGraphGenerator(g.Logger)
	transformCtx := transformers.NewTransformContext(g.BaseURL, g.OutputDirectory, notesRepository)

	pipeline := transformers.NewPipeline(
		transformers.NewSyntaxHighlighter(true),
		transformers.NewImageProcessor(g.Regexes.ObsidianImageRegex),
		transformers.NewHashtagEnricher(g.Regexes.HashtagRegex),
		transformers.NewWikilinkTransformer(
			g.Regexes.WikilinkRegex,
			g.Regexes.ObsidianImageRegex,
			g.Regexes.HashtagRegex,
			g.Logger,
			g.MarkdownRenderer,
		),
	)

	parseTime := g.parseNotes(scannedNotes, notesRepository)
	notesByPath := notesRepository.GetAllByPath()
	notesSlice := notesRepository.GetAllNotesSlice()
	progressManager.SendStepComplete(1, parseTime, len(notesSlice), nil)

	// Step 3: Build graph
	stepStart = time.Now()
	graphTime, graph := graphGenerator.Generate(notesByPath, notesRepository)
	progressManager.SendStepComplete(2, graphTime, len(notesByPath), nil)

	// Step 4: Transform content
	stepStart = time.Now()
	transformTime, _ := g.applyTransformationsConcurrently(pipeline, transformCtx, notesSlice)
	progressManager.SendStepComplete(3, transformTime, len(notesSlice), nil)

	// Step 5: Build folder tree
	stepStart = time.Now()
	buildTreeTime, fileTree, numberOfFolders := g.buildFolderTree(notesByPath)
	progressManager.SendStepComplete(4, buildTreeTime, numberOfFolders, nil)

	// Step 6: Execute templates
	stepStart = time.Now()
	sortedNotes := g.sortNotes(notesSlice)
	templateExecTime, numberOfTags, err := g.executeTemplatesBatch(sortedNotes, fileTree, graph)
	if err != nil {
		progressManager.SendStepComplete(5, time.Since(stepStart), 0, err)
		return err
	}
	progressManager.SendStepComplete(5, templateExecTime, len(sortedNotes), nil)

	// Step 7: Copy assets
	stepStart = time.Now()
	assetCopyTime := g.copyAssets(notesByPath)
	progressManager.SendStepComplete(6, assetCopyTime, len(notesByPath), nil)

	time.Sleep(100 * time.Millisecond)
	progressManager.Quit()

	if g.shouldPrintVaultHealth {
		g.printDiagnosticsSummary(notesSlice)
	}

	g.Logger.Debug("Site generation complete!",
		"duration", time.Since(start),
		"notes", len(notesByPath),
		"notes/second", float64(len(notesByPath))/time.Since(start).Seconds(),
		"avg ms/note", time.Since(start).Milliseconds()/int64(len(notesByPath)),
		"tags", numberOfTags,
		"preview pages", len(notesByPath),
		"folders", numberOfFolders,
	)
	return nil
}

type parseResult struct {
	note *models.ParsedNote
	err  error
}

func (g *StaticSiteGenerator) parseNotes(
	shallowNotes []*models.ScannedNote,
	notesRepository *repository.NoteRepository,
) time.Duration {
	start := time.Now()
	numWorkers := min(runtime.NumCPU(), len(shallowNotes))
	notesChan := make(chan *models.ScannedNote, len(shallowNotes))
	resultsChan := make(chan parseResult, len(shallowNotes))

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for note := range notesChan {
				result := g.parseNoteWorker(note)
				resultsChan <- result
			}
		}()
	}

	for _, note := range shallowNotes {
		notesChan <- note
	}
	close(notesChan)

	wg.Wait()
	close(resultsChan)

	for result := range resultsChan {
		if result.err != nil {
			continue
		}

		notesRepository.AddNote(result.note)
	}

	return time.Since(start)
}

func (g *StaticSiteGenerator) parseNoteWorker(
	shallowNote *models.ScannedNote,
) parseResult {
	if shallowNote == nil {
		return parseResult{
			note: nil,
			err:  fmt.Errorf("received nil shallowNote"),
		}
	}

	parsedNoteFromParser, err := g.Parser.ParseNote(shallowNote.FullPath, shallowNote.FileName)

	// passing as pointer is safe since each parseNoteWorker gets a different note
	parsedNote := &models.ParsedNote{
		ScannedNote: shallowNote,
	}

	if err != nil {
		return parseResult{
			note: parsedNote,
			err:  err,
		}
	}

	URL := strings.TrimPrefix(shallowNote.FullPath, g.InputDirectory)
	URL = strings.TrimPrefix(URL, "/")
	URL = strings.TrimSuffix(URL, ".md")

	slugifiedURL := "/" + utils.Slugify(URL) + ".html"
	parsedNote.URL = slugifiedURL

	parsedNote.Title = parsedNoteFromParser.Title
	parsedNote.Author = parsedNoteFromParser.Author
	parsedNote.RawBody = parsedNoteFromParser.RawBody
	parsedNote.Date = parsedNoteFromParser.Date
	parsedNote.UpdatedAt = parsedNoteFromParser.UpdatedAt
	parsedNote.Tags = parsedNoteFromParser.Tags
	parsedNote.Images = parsedNoteFromParser.Images
	parsedNote.Wikilinks = parsedNoteFromParser.Wikilinks
	parsedNote.Warnings = parsedNoteFromParser.Warnings
	parsedNote.MissingFiles = parsedNoteFromParser.MissingFiles

	return parseResult{
		note: parsedNote,
		err:  nil,
	}
}

func (g *StaticSiteGenerator) applyTransformationsConcurrently(
	pipeline *transformers.Pipeline,
	ctx *transformers.TransformContext,
	notes []*models.ParsedNote,
) (time.Duration, error) {
	start := time.Now()

	numWorkers := min(runtime.NumCPU(), len(notes))
	notesChan := make(chan *models.ParsedNote, len(notes))
	errorsChan := make(chan error, len(notes))

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for note := range notesChan {
				noteCtx := ctx.Clone()
				postBody := string(note.RawBody)

				transformedContent, err := pipeline.Transform(postBody, note, noteCtx)
				if err != nil {
					g.Logger.Warnf("failure to transform: %v", err)
				}

				finalHTMLBytes := markdown.ToHTML([]byte(transformedContent), nil, g.MarkdownRenderer)

				for _, embeddedPost := range noteCtx.EmbeddedPosts {
					finalHTMLBytes = bytes.Replace(finalHTMLBytes, []byte(embeddedPost.ID), []byte(embeddedPost.Content), 1)
				}

				note.HTMLContent = template.HTML(finalHTMLBytes)
				if err != nil {
					errorsChan <- err
				}
			}
		}()
	}

	for _, note := range notes {
		notesChan <- note
	}
	close(notesChan)

	wg.Wait()
	close(errorsChan)

	var errors []error
	for err := range errorsChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		g.Logger.Warn("some transformations failed", "count", len(errors))
		for _, err := range errors {
			g.Logger.Error("transformation error", "error", err)
		}
	}

	return time.Since(start), nil
}

func (g *StaticSiteGenerator) buildFolderTree(
	notes map[string]*models.ParsedNote,
) (time.Duration, *models.Folder, int) {
	start := time.Now()
	root := &models.Folder{
		Name:     "Home",
		Path:     "",
		Posts:    []*models.ParsedNote{},
		Children: make(map[string]*models.Folder),
	}
	numberOfFolders := 0

	for _, note := range notes {
		if !note.IsInsideFolder {
			root.Posts = append(root.Posts, note)
			continue
		}

		htmlFileName := utils.Slugify(strings.TrimSuffix(note.FileName, ".md")) + ".html"
		url := strings.TrimRight(note.URL, htmlFileName)
		parts := strings.Split(url, string(filepath.Separator))
		currentNode := root

		for _, part := range parts {
			if part == "" {
				continue
			}
			lowercasePart := strings.ToLower(part)

			if _, ok := currentNode.Children[lowercasePart]; !ok {
				childPath := filepath.ToSlash(filepath.Join(currentNode.Path, lowercasePart))

				currentNode.Children[lowercasePart] = &models.Folder{
					Name:     lowercasePart,
					Path:     childPath,
					Posts:    []*models.ParsedNote{},
					Children: make(map[string]*models.Folder),
				}
				numberOfFolders++
			}
			currentNode = currentNode.Children[lowercasePart]
		}

		currentNode.Posts = append(currentNode.Posts, note)
	}

	sortFolderTree(root)
	return time.Since(start), root, numberOfFolders
}

func sortFolderTree(folder *models.Folder) {
	sort.Slice(folder.Posts, func(i, j int) bool {
		return folder.Posts[i].Title < folder.Posts[j].Title
	})

	for _, child := range folder.Children {
		sortFolderTree(child)
	}
}

func (g *StaticSiteGenerator) sortNotes(notes []*models.ParsedNote) []*models.ParsedNote {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Date.After(notes[j].Date)
	})
	return notes
}

func (g *StaticSiteGenerator) executeTemplatesBatch(
	notes []*models.ParsedNote,
	fileTree *models.Folder,
	graph []byte,
) (time.Duration, int, error) {
	start := time.Now()
	batchWriter := executors.NewBatchWriter(g.OutputDirectory, g.Logger, g.Templates)

	tags, postsByTag := g.computeTags(notes)
	baseData := g.createBaseTemplateData(fileTree, graph, tags)

	for _, note := range notes {
		data := struct {
			models.ParsedNote
			SiteTitle    string
			SiteSubtitle string
			BaseURL      string
			CurrentYear  int
			FileTree     *models.Folder
			Graph        template.JS
			Tags         []models.Tag
		}{
			ParsedNote:   *note,
			SiteTitle:    baseData.SiteTitle,
			SiteSubtitle: baseData.SiteSubtitle,
			BaseURL:      baseData.BaseURL,
			CurrentYear:  baseData.CurrentYear,
			FileTree:     baseData.FileTree,
			Graph:        baseData.Graph,
			Tags:         baseData.Tags,
		}

		filePath := note.URL
		if !strings.HasSuffix(filePath, ".html") {
			filePath += ".html"
		}
		cleanFilePath := strings.TrimPrefix(filePath, "/")

		batchWriter.AddFile(cleanFilePath, "post.html", data)

		previewPath := filepath.Join("previews", cleanFilePath)
		batchWriter.AddFile(previewPath, "preview.html", data)
	}

	for _, tag := range tags {
		if err := g.paginateTagPages(batchWriter, tag, postsByTag[tag.Slug], baseData); err != nil {
			return time.Since(start), 0, err
		}
	}

	if err := g.paginateIndexPages(batchWriter, notes, baseData); err != nil {
		return time.Since(start), 0, err
	}

	if err := g.addFolderPagesToBatchRecursive(batchWriter, fileTree, baseData); err != nil {
		return time.Since(start), 0, err
	}

	batchWriter.AddFile("404.html", "404.html", baseData)

	g.Logger.Debug("Writing all template files")
	if err := batchWriter.WriteAllFiles(); err != nil {
		return time.Since(start), 0, fmt.Errorf("failed to write template files: %w", err)
	}

	return time.Since(start), len(tags), nil
}

func (g *StaticSiteGenerator) addFolderPagesToBatchRecursive(
	batchWriter *executors.BatchWriter,
	folder *models.Folder,
	baseData BaseTemplateData,
) error {
	if folder.Path != "" {
		if err := g.paginateFolderPages(batchWriter, folder, baseData); err != nil {
			return err
		}
	}

	for _, child := range folder.Children {
		if err := g.addFolderPagesToBatchRecursive(batchWriter, child, baseData); err != nil {
			return err
		}
	}

	return nil
}

func (g *StaticSiteGenerator) paginateIndexPages(
	batchWriter *executors.BatchWriter,
	posts []*models.ParsedNote,
	baseData BaseTemplateData,
) error {
	if len(posts) == 0 {
		// Create empty index page
		data := struct {
			BaseTemplateData
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Posts:            posts,
			Pagination:       PaginationData{},
		}
		return batchWriter.AddFile("index.html", "index.html", data)
	}

	totalPages := (len(posts) + g.NotesPerPage - 1) / g.NotesPerPage

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * g.NotesPerPage
		end := min(start+g.NotesPerPage, len(posts))
		pagePosts := posts[start:end]

		var pagePath string
		if page == 1 {
			pagePath = "index.html"
		} else {
			pagePath = filepath.Join("page", fmt.Sprintf("%d", page), "index.html")
		}

		pagination := PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		}
		if page > 1 {
			pagination.HasPrev = true
			if page == 2 {
				pagination.PrevPageURL = g.BaseURL
			} else {
				pagination.PrevPageURL = g.BaseURL + fmt.Sprintf("page/%d/", page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = g.BaseURL + fmt.Sprintf("page/%d/", page+1)
		}

		data := struct {
			BaseTemplateData
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Posts:            pagePosts,
			Pagination:       pagination,
		}

		if err := batchWriter.AddFile(pagePath, "index.html", data); err != nil {
			return fmt.Errorf("failed to add index page %d: %w", page, err)
		}
	}

	return nil
}

func (g *StaticSiteGenerator) paginateTagPages(
	batchWriter *executors.BatchWriter,
	tag models.Tag,
	posts []*models.ParsedNote,
	baseData BaseTemplateData,
) error {
	if len(posts) == 0 {
		return nil
	}

	totalPages := (len(posts) + g.NotesPerPage - 1) / g.NotesPerPage
	pathPrefix := filepath.Join("tag", tag.Slug)

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * g.NotesPerPage
		end := min(start+g.NotesPerPage, len(posts))
		pagePosts := posts[start:end]

		var pagePath string
		if page == 1 {
			pagePath = filepath.Join(pathPrefix, "index.html")
		} else {
			pagePath = filepath.Join(pathPrefix, "page", fmt.Sprintf("%d", page), "index.html")
		}

		pagination := PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		}
		if page > 1 {
			pagination.HasPrev = true
			if page == 2 {
				pagination.PrevPageURL = g.BaseURL + strings.ReplaceAll(pathPrefix, "\\", "/") + "/"
			} else {
				pagination.PrevPageURL = g.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(pathPrefix, "\\", "/"), page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = g.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(pathPrefix, "\\", "/"), page+1)
		}

		data := struct {
			BaseTemplateData
			models.Tag
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Tag:              tag,
			Posts:            pagePosts,
			Pagination:       pagination,
		}

		if err := batchWriter.AddFile(pagePath, "tag.html", data); err != nil {
			return fmt.Errorf("failed to add tag page %s page %d: %w", tag.Slug, page, err)
		}
	}

	return nil
}

func (g *StaticSiteGenerator) paginateFolderPages(
	batchWriter *executors.BatchWriter,
	folder *models.Folder,
	baseData BaseTemplateData,
) error {
	if len(folder.Posts) == 0 {
		data := struct {
			BaseTemplateData
			Folder     *models.Folder
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Folder:           folder,
			Posts:            folder.Posts,
			Pagination:       PaginationData{},
		}
		folderPath := filepath.Join(folder.Path, "index.html")
		return batchWriter.AddFile(folderPath, "folder.html", data)
	}

	totalPages := (len(folder.Posts) + g.NotesPerPage - 1) / g.NotesPerPage

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * g.NotesPerPage
		end := min(start+g.NotesPerPage, len(folder.Posts))
		pagePosts := folder.Posts[start:end]

		var pagePath string
		if page == 1 {
			pagePath = filepath.Join(folder.Path, "index.html")
		} else {
			pagePath = filepath.Join(folder.Path, "page", fmt.Sprintf("%d", page), "index.html")
		}

		pagination := PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		}
		if page > 1 {
			pagination.HasPrev = true
			if page == 2 {
				pagination.PrevPageURL = g.BaseURL + strings.ReplaceAll(folder.Path, "\\", "/") + "/"
			} else {
				pagination.PrevPageURL = g.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(folder.Path, "\\", "/"), page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = g.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(folder.Path, "\\", "/"), page+1)
		}

		data := struct {
			BaseTemplateData
			Folder     *models.Folder
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Folder:           folder,
			Posts:            pagePosts,
			Pagination:       pagination,
		}

		if err := batchWriter.AddFile(pagePath, "folder.html", data); err != nil {
			return fmt.Errorf("failed to add folder page %s page %d: %w", folder.Name, page, err)
		}
	}

	return nil
}

func (g *StaticSiteGenerator) copyAssets(notesByPath map[string]*models.ParsedNote) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	errorChan := make(chan error, 100)

	fullImagePaths := make(map[string]string)
	for _, p := range notesByPath {
		for _, img := range p.Images {
			fullPath := filepath.Join(filepath.Dir(p.FullPath), img.RelativePath)
			fullImagePaths[img.RelativePath] = fullPath
		}
	}

	g.Logger.Debug("Copying images 🖼️ ", "type", "images", "count", len(fullImagePaths))
	for name, path := range fullImagePaths {
		wg.Add(1)
		go func(name, path string) {
			defer wg.Done()
			destPath := filepath.Join(g.OutputDirectory, "images", name)
			if err := utils.CopyFile(path, destPath); err != nil {
				errorChan <- fmt.Errorf("image %s: %w", name, err)
			}
		}(name, path)
	}

	for _, dir := range []string{"js", "css", "assets"} {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			g.Logger.Debug("Copying static assets 📦", "dir", dir)
			if err := utils.CopyStaticDirectory(dir, g.OutputDirectory); err != nil {
				errorChan <- fmt.Errorf("asset folder '%s': %w", dir, err)
			}
		}(dir)
	}

	wg.Wait()
	close(errorChan)

	for err := range errorChan {
		g.Logger.Debugf("%s\n", err.Error())
	}

	return time.Since(start)
}

func (g *StaticSiteGenerator) createBaseTemplateData(
	fileTree *models.Folder,
	graph []byte,
	tags []models.Tag,
) BaseTemplateData {
	return BaseTemplateData{
		SiteTitle:    g.SiteTitle,
		SiteSubtitle: g.SiteSubtitle,
		BaseURL:      g.BaseURL,
		CurrentYear:  time.Now().Year(),
		FileTree:     fileTree,
		Graph:        template.JS(graph),
		Tags:         tags,
	}
}

func (g *StaticSiteGenerator) computeTags(
	notes []*models.ParsedNote,
) ([]models.Tag, map[string][]*models.ParsedNote) {
	tagsMap := make(map[string]*models.Tag)
	for _, note := range notes {
		for _, tag := range note.Tags {
			tagsMap[tag.Slug] = &tag
		}
	}

	tags := make([]models.Tag, 0, len(tagsMap))
	for _, tag := range tagsMap {
		tags = append(tags, *tag)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})

	postsByTag := make(map[string][]*models.ParsedNote, len(tags))

	for _, note := range notes {
		for _, tag := range note.Tags {
			postsByTag[tag.Slug] = append(postsByTag[tag.Slug], note)
		}
	}

	return tags, postsByTag
}

func (g *StaticSiteGenerator) printDiagnosticsSummary(notes []*models.ParsedNote) {
	var totalWarnings, totalMissingFiles int
	var notesWithWarnings, notesWithMissingFiles int
	var warningTypes = make(map[string]int)
	var missingFileTypes = make(map[string]int)

	for _, note := range notes {
		if note.HasWarnings() {
			notesWithWarnings++
			totalWarnings += len(note.Warnings)

			for _, warning := range note.Warnings {
				if strings.Contains(warning, "unknown frontmatter field") {
					warningTypes["unknown_frontmatter"]++
				} else if strings.Contains(warning, "frontmatter parsing issues") {
					warningTypes["frontmatter_parsing"]++
				} else {
					warningTypes["other"]++
				}
			}
		}

		if note.HasMissingFiles() {
			notesWithMissingFiles++
			totalMissingFiles += len(note.MissingFiles)

			for _, file := range note.MissingFiles {
				ext := strings.ToLower(filepath.Ext(file))
				if ext == "" {
					ext = "no_extension"
				}
				missingFileTypes[ext]++
			}
		}
	}

	summary := terminal.DiagnosticsSummary{
		TotalNotes:            len(notes),
		NotesWithWarnings:     notesWithWarnings,
		NotesWithMissingFiles: notesWithMissingFiles,
		TotalWarnings:         totalWarnings,
		TotalMissingFiles:     totalMissingFiles,
		WarningTypes:          warningTypes,
		MissingFileTypes:      missingFileTypes,
	}

	formatter := terminal.NewDiagnosticsFormatter()
	output := formatter.FormatDiagnosticsSummary(summary)
	fmt.Print(output)

	if totalWarnings > 0 || totalMissingFiles > 0 {
		g.exportDiagnostics(notes, warningTypes, missingFileTypes, totalWarnings, totalMissingFiles)
		g.Logger.Debug("📊 Diagnostics report exported", "path", filepath.Join(g.OutputDirectory, "vault-diagnostics.json"), "size", "0.5 KB")
	}
}

func (g *StaticSiteGenerator) exportDiagnostics(
	notes []*models.ParsedNote,
	warningTypes map[string]int,
	missingFileTypes map[string]int,
	totalWarnings int,
	totalMissingFiles int,
) {
	type NoteDiagnostic struct {
		Title        string   `json:"title"`
		RelativePath string   `json:"relative_path"`
		Warnings     []string `json:"warnings"`
		MissingFiles []string `json:"missing_files"`
		TotalIssues  int      `json:"total_issues"`
	}

	type DiagnosticReport struct {
		GeneratedAt           string           `json:"generated_at"`
		TotalNotes            int              `json:"total_notes"`
		NotesWithWarnings     int              `json:"notes_with_warnings"`
		NotesWithMissingFiles int              `json:"notes_with_missing_files"`
		TotalWarnings         int              `json:"total_warnings"`
		TotalMissingFiles     int              `json:"total_missing_files"`
		WarningTypes          map[string]int   `json:"warning_types"`
		MissingFileTypes      map[string]int   `json:"missing_file_types"`
		ProblematicNotes      []NoteDiagnostic `json:"problematic_notes"`
	}

	if totalWarnings == 0 && totalMissingFiles == 0 {
		return
	}

	var problematicNotes []NoteDiagnostic
	var notesWithWarnings, notesWithMissingFiles int

	for _, note := range notes {
		if note.HasWarnings() || note.HasMissingFiles() {
			totalIssues := len(note.Warnings) + len(note.MissingFiles)
			problematicNotes = append(problematicNotes, NoteDiagnostic{
				Title:        note.Title,
				RelativePath: note.RelativePath,
				Warnings:     note.Warnings,
				MissingFiles: note.MissingFiles,
				TotalIssues:  totalIssues,
			})

			if note.HasWarnings() {
				notesWithWarnings++
			}
			if note.HasMissingFiles() {
				notesWithMissingFiles++
			}
		}
	}

	report := DiagnosticReport{
		GeneratedAt:           time.Now().Format(time.RFC3339),
		TotalNotes:            len(notes),
		NotesWithWarnings:     notesWithWarnings,
		NotesWithMissingFiles: notesWithMissingFiles,
		TotalWarnings:         totalWarnings,
		TotalMissingFiles:     totalMissingFiles,
		WarningTypes:          warningTypes,
		MissingFileTypes:      missingFileTypes,
		ProblematicNotes:      problematicNotes,
	}

	diagnosticsPath := filepath.Join(g.OutputDirectory, "vault-diagnostics.json")
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		g.Logger.Warn("Failed to marshal diagnostics report", "error", err)
		return
	}

	if err := os.WriteFile(diagnosticsPath, jsonData, 0644); err != nil {
		g.Logger.Warn("Failed to write diagnostics report", "error", err)
		return
	}

	g.Logger.Debug("📊 Diagnostics report exported",
		"path", diagnosticsPath,
		"size", fmt.Sprintf("%.1f KB", float64(len(jsonData))/1024),
	)
}
