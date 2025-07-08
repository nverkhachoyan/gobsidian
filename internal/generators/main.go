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

	"regexp"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"

	"gobsidian/internal/config"
	"gobsidian/internal/executors"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/repository"
	"gobsidian/internal/scanner"
	"gobsidian/internal/terminal"
	"gobsidian/internal/transformers"
	"gobsidian/internal/utils"

	"github.com/gomarkdown/markdown/html"

	"github.com/gomarkdown/markdown"
)

type Generator interface {
	Generate(scannedNotes []*models.ScannedNote) error
}

type StaticSiteGenerator struct {
	SiteConfig             config.SiteConfig
	Regexes                *config.Regexes
	Logger                 *log.Logger
	Templates              *template.Template
	Parser                 parsers.Parser
	Scanner                *scanner.NoteScanner
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
	FileTreeJSON template.JS
}

func NewStaticSiteGenerator(
	siteConfig config.SiteConfig,
	regexes *config.Regexes,
	logger *log.Logger,
	templates *template.Template,
	parser parsers.Parser,
	scanner *scanner.NoteScanner,
	markdownRenderer *html.Renderer,
	shouldPrintVaultHealth bool,
) (*StaticSiteGenerator, error) {
	return &StaticSiteGenerator{
		SiteConfig:             siteConfig,
		Regexes:                regexes,
		Logger:                 logger,
		Templates:              templates,
		Parser:                 parser,
		Scanner:                scanner,
		MarkdownRenderer:       markdownRenderer,
		shouldPrintVaultHealth: shouldPrintVaultHealth,
	}, nil
}

func (g *StaticSiteGenerator) Generate(scannedNotes []*models.ScannedNote) error {
	start := time.Now()

	printer := terminal.NewPrettyPrinter()
	reporter := terminal.NewProgressReporter(printer)

	parseStep := terminal.Step{Name: "Parsing notes", Icon: "📝"}

	notesRepository := repository.NewNoteRepository()
	graphGenerator := NewGraphGenerator(g.Logger)
	transformCtx := transformers.NewTransformContext(g.SiteConfig.BaseURL, g.SiteConfig.OutputDirectory, notesRepository)

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

	reporter.ReportStep(parseStep, terminal.StepResult{
		Duration: parseTime,
		Count:    len(notesSlice),
		Error:    nil,
	})

	graphStep := terminal.Step{Name: "Built graph", Icon: "🔗"}
	graphTime, graphJSON := graphGenerator.Generate(notesByPath, notesRepository)

	reporter.ReportStep(graphStep, terminal.StepResult{
		Duration: graphTime,
		Count:    len(notesByPath),
		Error:    nil,
	})

	transformStep := terminal.Step{Name: "Transformed content", Icon: "✨"}
	transformTime, _ := g.applyTransformationsConcurrently(pipeline, transformCtx, notesSlice)

	reporter.ReportStep(transformStep, terminal.StepResult{
		Duration: transformTime,
		Count:    len(notesSlice),
		Error:    nil,
	})

	treeStep := terminal.Step{Name: "Built folder tree", Icon: "📁"}
	buildTreeTime, fileTree, numberOfFolders, fileTreeJSON, err := g.buildFolderTree(notesByPath)

	if err != nil {
		reporter.ReportStep(treeStep, terminal.StepResult{
			Duration: buildTreeTime,
			Count:    numberOfFolders,
			Error:    err,
		})
		return fmt.Errorf("failed to build folder tree: %w", err)
	}

	reporter.ReportStep(treeStep, terminal.StepResult{
		Duration: buildTreeTime,
		Count:    numberOfFolders,
		Error:    nil,
	})

	templateStep := terminal.Step{Name: "Executed templates", Icon: "🎨"}
	templateStart := time.Now()
	sortedNotes := g.sortNotes(notesSlice)
	templateExecTime, numberOfTags, err := g.executeTemplatesBatch(sortedNotes, fileTree, graphJSON, fileTreeJSON)
	if err != nil {
		reporter.ReportStep(templateStep, terminal.StepResult{
			Duration: time.Since(templateStart),
			Count:    0,
			Error:    err,
		})
		return err
	}

	reporter.ReportStep(templateStep, terminal.StepResult{
		Duration: templateExecTime,
		Count:    len(sortedNotes),
		Error:    nil,
	})

	assetsStep := terminal.Step{Name: "Copied assets", Icon: "📦"}
	assetCopyTime := g.copyAssets(notesByPath, graphJSON, fileTreeJSON)
	g.generateSyntaxHighlighterCSS()

	reporter.ReportStep(assetsStep, terminal.StepResult{
		Duration: assetCopyTime,
		Count:    len(notesByPath),
		Error:    nil,
	})

	if g.shouldPrintVaultHealth {
		g.printDiagnosticsSummary(notesSlice)
	}

	g.Logger.Debug("Site generation complete!",
		"duration", time.Since(start),
		"notes", len(notesByPath),
		"notes/second", float64(len(notesByPath))/time.Since(start).Seconds(),
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

	URL := strings.TrimPrefix(shallowNote.FullPath, g.SiteConfig.InputDirectory)
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

// File represents a note within the JSON file tree.
type File struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// FolderNode represents a folder in the JSON file tree.
type FolderNode struct {
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Children []*FolderNode `json:"children"`
	Files    []*File       `json:"files"`
}

func convertFolderToNode(folder *models.Folder) *FolderNode {
	node := &FolderNode{
		Name:     folder.Name,
		Path:     folder.Path,
		Children: make([]*FolderNode, 0),
		Files:    make([]*File, 0),
	}

	for _, post := range folder.Posts {
		node.Files = append(node.Files, &File{
			Name: post.Title,
			URL:  post.URL,
		})
	}
	sort.Slice(node.Files, func(i, j int) bool {
		return node.Files[i].Name < node.Files[j].Name
	})

	childKeys := make([]string, 0, len(folder.Children))
	for key := range folder.Children {
		childKeys = append(childKeys, key)
	}
	sort.Strings(childKeys)

	for _, key := range childKeys {
		childFolder := folder.Children[key]
		node.Children = append(node.Children, convertFolderToNode(childFolder))
	}

	return node
}

func (g *StaticSiteGenerator) buildFolderTree(
	notes map[string]*models.ParsedNote,
) (time.Duration, *models.Folder, int, []byte, error) {
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

	jsonTree := convertFolderToNode(root)
	jsonBytes, err := json.Marshal(jsonTree)
	if err != nil {
		return time.Since(start), nil, 0, nil, fmt.Errorf("failed to marshal folder tree to JSON: %w", err)
	}

	return time.Since(start), root, numberOfFolders, jsonBytes, nil
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
	fileTreeJSON []byte,
) (time.Duration, int, error) {
	start := time.Now()
	batchWriter := executors.NewBatchWriter(g.SiteConfig.OutputDirectory, g.Logger, g.Templates)

	tags, postsByTag := g.computeTags(notes)
	baseData := g.createBaseTemplateData(fileTree, graph, tags, fileTreeJSON)

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
			FileTreeJSON template.JS
		}{
			ParsedNote:   *note,
			SiteTitle:    baseData.SiteTitle,
			SiteSubtitle: baseData.SiteSubtitle,
			BaseURL:      baseData.BaseURL,
			CurrentYear:  baseData.CurrentYear,
			FileTree:     baseData.FileTree,
			Graph:        baseData.Graph,
			Tags:         baseData.Tags,
			FileTreeJSON: baseData.FileTreeJSON,
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

	totalPages := (len(posts) + g.SiteConfig.NotesPerPage - 1) / g.SiteConfig.NotesPerPage

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * g.SiteConfig.NotesPerPage
		end := min(start+g.SiteConfig.NotesPerPage, len(posts))
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
				pagination.PrevPageURL = g.SiteConfig.BaseURL
			} else {
				pagination.PrevPageURL = g.SiteConfig.BaseURL + fmt.Sprintf("page/%d/", page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = g.SiteConfig.BaseURL + fmt.Sprintf("page/%d/", page+1)
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

	totalPages := (len(posts) + g.SiteConfig.NotesPerPage - 1) / g.SiteConfig.NotesPerPage
	pathPrefix := filepath.Join("tag", tag.Slug)

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * g.SiteConfig.NotesPerPage
		end := min(start+g.SiteConfig.NotesPerPage, len(posts))
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
				pagination.PrevPageURL = g.SiteConfig.BaseURL + strings.ReplaceAll(pathPrefix, "\\", "/") + "/"
			} else {
				pagination.PrevPageURL = g.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(pathPrefix, "\\", "/"), page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = g.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(pathPrefix, "\\", "/"), page+1)
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

	totalPages := (len(folder.Posts) + g.SiteConfig.NotesPerPage - 1) / g.SiteConfig.NotesPerPage

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * g.SiteConfig.NotesPerPage
		end := min(start+g.SiteConfig.NotesPerPage, len(folder.Posts))
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
				pagination.PrevPageURL = g.SiteConfig.BaseURL + strings.ReplaceAll(folder.Path, "\\", "/") + "/"
			} else {
				pagination.PrevPageURL = g.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(folder.Path, "\\", "/"), page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = g.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(folder.Path, "\\", "/"), page+1)
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

func (g *StaticSiteGenerator) copyAssets(notesByPath map[string]*models.ParsedNote, graph []byte, fileTreeJSON []byte) time.Duration {
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
			destPath := filepath.Join(g.SiteConfig.OutputDirectory, "images", name)
			if err := utils.CopyFile(path, destPath); err != nil {
				errorChan <- fmt.Errorf("image %s: %w", name, err)
			}
		}(name, path)
	}

	if len(fileTreeJSON) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			destPath := filepath.Join(g.SiteConfig.OutputDirectory, "file-tree.json")
			g.Logger.Debug("Writing file-tree.json 🌲")
			if err := os.WriteFile(destPath, fileTreeJSON, 0644); err != nil {
				errorChan <- fmt.Errorf("file-tree.json: %w", err)
			}
		}()
	}

	for _, dir := range []string{"assets", "css", "js"} {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			g.Logger.Debug("Copying static assets 📦", "dir", dir)
			if err := utils.CopyStaticDirectory(dir, g.SiteConfig.OutputDirectory); err != nil {
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
	fileTreeJSON []byte,
) BaseTemplateData {
	return BaseTemplateData{
		SiteTitle:    g.SiteConfig.SiteTitle,
		SiteSubtitle: g.SiteConfig.SiteSubtitle,
		BaseURL:      g.SiteConfig.BaseURL,
		CurrentYear:  time.Now().Year(),
		FileTree:     fileTree,
		Graph:        template.JS(graph),
		Tags:         tags,
		FileTreeJSON: template.JS(fileTreeJSON),
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

	printer := terminal.NewPrettyPrinter()
	output := printer.PrintDiagnostics(summary)
	fmt.Println(printer.Box("Vault diagnostics", output))

	if totalWarnings > 0 || totalMissingFiles > 0 {
		g.exportDiagnostics(notes, warningTypes, missingFileTypes, totalWarnings, totalMissingFiles)
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

	diagnosticsPath := filepath.Join(g.SiteConfig.OutputDirectory, "vault-diagnostics.json")
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

func (g *StaticSiteGenerator) generateSyntaxHighlighterCSS() error {
	var lightBuf, darkBuf bytes.Buffer
	formatter := chromahtml.New(chromahtml.WithClasses(true))

	if err := formatter.WriteCSS(
		&lightBuf,
		styles.Get(g.SiteConfig.Theme.SyntaxHighlighterLight),
	); err != nil {
		return fmt.Errorf("failed to generate light theme: %w", err)
	}

	if err := formatter.WriteCSS(
		&darkBuf,
		styles.Get(g.SiteConfig.Theme.SyntaxHighlighterDark),
	); err != nil {
		return fmt.Errorf("failed to generate dark theme: %w", err)
	}

	scopedDarkCSS := g.scopeCSS(darkBuf.String(), "html.dark-mode")

	// combine light and dark themes
	finalCSS := lightBuf.String() + "\n" + scopedDarkCSS
	outputPath := filepath.Join(g.SiteConfig.OutputDirectory, "assets", "chroma.css")

	if err := os.WriteFile(outputPath, []byte(finalCSS), 0644); err != nil {
		return fmt.Errorf("failed to write combined CSS file: %w", err)
	}

	g.Logger.Debug("Successfully generated combined Chroma CSS at %s", outputPath)
	return nil
}

// util function to scope CSS to a specific selector
func (g *StaticSiteGenerator) scopeCSS(css, scope string) string {
	var result strings.Builder
	r := regexp.MustCompile(`(?s)([^}]+)({[^}]+})`)
	submatches := r.FindAllStringSubmatch(css, -1)

	for _, submatch := range submatches {
		selectors := strings.TrimSpace(submatch[1])
		declarations := submatch[2]

		if strings.HasPrefix(selectors, "@") {
			result.WriteString(selectors)
			result.WriteString(declarations)
			continue
		}

		var scopedSelectors []string
		for _, selector := range strings.Split(selectors, ",") {
			scopedSelectors = append(scopedSelectors, scope+" "+strings.TrimSpace(selector))
		}
		result.WriteString(strings.Join(scopedSelectors, ", "))
		result.WriteString(declarations)
	}
	return result.String()
}
