package generators

import (
	"bytes"
	"fmt"
	"html/template"
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
	"gobsidian/internal/transformers"
	"gobsidian/internal/utils"

	"github.com/gomarkdown/markdown/html"

	"github.com/gomarkdown/markdown"
)

type StaticSiteGenerator struct {
	InputDirectory   string
	OutputDirectory  string
	SiteTitle        string
	SiteSubtitle     string
	BaseURL          string
	Regexes          *models.ParserRegexes
	Logger           *log.Logger
	Templates        *template.Template
	Parser           parsers.Parser
	MarkdownRenderer *html.Renderer
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

func NewStaticSiteGenerator(cfg *StaticSiteGenerator) (*StaticSiteGenerator, error) {
	return &StaticSiteGenerator{
		InputDirectory:   cfg.InputDirectory,
		OutputDirectory:  cfg.OutputDirectory,
		SiteTitle:        cfg.SiteTitle,
		SiteSubtitle:     cfg.SiteSubtitle,
		BaseURL:          cfg.BaseURL,
		Regexes:          cfg.Regexes,
		Logger:           cfg.Logger,
		Templates:        cfg.Templates,
		Parser:           cfg.Parser,
		MarkdownRenderer: cfg.MarkdownRenderer,
	}, nil
}

func (g *StaticSiteGenerator) GenerateSite() error {
	start := time.Now()
	g.Logger.Print("Generating...")

	noteScanner := NewNoteScanner(g.Logger, g.InputDirectory)
	err := noteScanner.ScanAllNotes()
	if err != nil {
		return err
	}
	scannedNotes := noteScanner.GetAllNotes()

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

	g.parseNotes(scannedNotes, notesRepository)
	notesByPath := notesRepository.GetAllByPath()
	notesSlice := notesRepository.GetAllNotesSlice()

	g.applyTransformationsConcurrently(pipeline, transformCtx, notesSlice)

	fileTree, numberOfFolders := g.buildFolderTree(notesByPath)
	graph := graphGenerator.Generate(notesByPath, notesRepository)

	numberOfTags, err := g.executeTemplatesBatch(notesSlice, fileTree, graph)
	if err != nil {
		return err
	}

	g.copyAssets(notesByPath)

	g.Logger.Print("Site generation complete!",
		"duration", time.Since(start),
		"notes", len(notesByPath),
		"notes_per_second", float64(len(notesByPath))/time.Since(start).Seconds(),
		"avg_ms_per_note", time.Since(start).Milliseconds()/int64(len(notesByPath)),
		"tags", numberOfTags,
		"index_pages", 1,
		"preview_pages", len(notesByPath),
		"folders", numberOfFolders,
	)
	return nil
}

func (g *StaticSiteGenerator) applyTransformationsConcurrently(pipeline *transformers.Pipeline, ctx *transformers.TransformContext, notes []*models.ParsedNote) error {
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

				postBody := string(note.RawBody)

				transformedContent, err := pipeline.Transform(postBody, note, ctx)
				if err != nil {
					g.Logger.Warnf("failure to transform: %v", err)
				}

				finalHTMLBytes := markdown.ToHTML([]byte(transformedContent), nil, g.MarkdownRenderer)

				for _, embeddedPost := range ctx.EmbeddedPosts {
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

	g.Logger.Printf("Transformation took %s", time.Since(start))
	return nil
}

func (g *StaticSiteGenerator) buildFolderTree(notes map[string]*models.ParsedNote) (*models.Folder, int) {
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
	g.Logger.Printf("Directory creation took %s", time.Since(start))
	return root, numberOfFolders
}

func sortFolderTree(folder *models.Folder) {
	// Sort posts in the current folder alphabetically by title
	sort.Slice(folder.Posts, func(i, j int) bool {
		return folder.Posts[i].Title < folder.Posts[j].Title
	})

	// Recurse into children
	for _, child := range folder.Children {
		sortFolderTree(child)
	}
}

func (g *StaticSiteGenerator) executeTemplatesBatch(notes []*models.ParsedNote, fileTree *models.Folder, graph []byte) (int, error) {
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

		// Make sure URL ends with .html
		filePath := note.URL
		if !strings.HasSuffix(filePath, ".html") {
			filePath += ".html"
		}

		// Remove leading slash if present
		filePath = strings.TrimPrefix(filePath, "/")
		batchWriter.AddFile(filePath, "post.html", data)

		previewPath := filepath.Join("previews", strings.TrimPrefix(note.URL, "/"))
		if !strings.HasSuffix(previewPath, ".html") {
			previewPath += ".html"
		}
		batchWriter.AddFile(previewPath, "preview.html", data)
	}

	for _, tag := range tags {
		data := struct {
			BaseTemplateData
			models.Tag
			Posts []*models.ParsedNote
		}{
			BaseTemplateData: baseData,
			Tag:              tag,
			Posts:            postsByTag[tag.Slug],
		}
		tagPath := filepath.Join("tag", tag.Slug, "index.html")
		batchWriter.AddFile(tagPath, "tag.html", data)
	}

	batchWriter.AddFile("index.html", "index.html", struct {
		BaseTemplateData
		Posts []*models.ParsedNote
	}{
		BaseTemplateData: baseData,
		Posts:            notes,
	})

	g.addFolderPagesToBatch(batchWriter, fileTree, baseData)

	batchWriter.AddFile("404.html", "404.html", baseData)

	g.Logger.Debug("Writing all template files")
	if err := batchWriter.WriteAllFiles(); err != nil {
		return 0, fmt.Errorf("failed to write template files: %w", err)
	}

	g.Logger.Printf("Finished writing template files in %s", time.Since(start))
	return len(tags), nil
}

func (g *StaticSiteGenerator) createBaseTemplateData(fileTree *models.Folder, graph []byte, tags []models.Tag) BaseTemplateData {
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

func (g *StaticSiteGenerator) addFolderPagesToBatch(
	batchWriter *executors.BatchWriter,
	folder *models.Folder,
	baseData BaseTemplateData,
) error {
	return g.addFolderPagesToBatchRecursive(batchWriter, folder, baseData)
}

func (g *StaticSiteGenerator) addFolderPagesToBatchRecursive(
	batchWriter *executors.BatchWriter,
	folder *models.Folder,
	baseData BaseTemplateData,
) error {
	if folder.Path != "" {
		folderPath := filepath.Join(folder.Path, "index.html")

		folderData := struct {
			BaseTemplateData
			Folder *models.Folder
		}{
			BaseTemplateData: baseData,
			Folder:           folder,
		}

		if err := batchWriter.AddFile(folderPath, "folder.html", folderData); err != nil {
			return fmt.Errorf("failed to add folder page for %s: %w", folder.Name, err)
		}
	}

	for _, child := range folder.Children {
		if err := g.addFolderPagesToBatchRecursive(batchWriter, child, baseData); err != nil {
			return err
		}
	}

	return nil
}

func (g *StaticSiteGenerator) computeTags(notes []*models.ParsedNote) ([]models.Tag, map[string][]*models.ParsedNote) {
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

type parseResult struct {
	note *models.ParsedNote
	err  error
}

func (g *StaticSiteGenerator) parseNotes(shallowNotes []*models.ScannedNote, notesRepository *repository.NoteRepository) {
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
			g.Logger.Warn("failed to parse note", "error", result.err)
			continue
		}

		notesRepository.AddNote(result.note)
	}

	g.Logger.Printf("Wikilink parsing took %s", time.Since(start))
}

func (g *StaticSiteGenerator) parseNoteWorker(shallowNote *models.ScannedNote) parseResult {
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
		g.Logger.Warn("failed to parse note", "path", shallowNote.FullPath, "error", err)
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

	return parseResult{
		note: parsedNote,
		err:  nil,
	}
}

func (g *StaticSiteGenerator) copyAssets(notesByPath map[string]*models.ParsedNote) {
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

	g.Logger.Print("Copying images 🖼️ ", "type", "images", "count", len(fullImagePaths))
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
			g.Logger.Print("Copying static assets 📦", "dir", dir)
			if err := utils.CopyStaticDirectory(dir, g.OutputDirectory); err != nil {
				errorChan <- fmt.Errorf("asset folder '%s': %w", dir, err)
			}
		}(dir)
	}

	wg.Wait()
	close(errorChan)

	var allErrors []string
	for err := range errorChan {
		allErrors = append(allErrors, err.Error())
	}

	if len(allErrors) > 0 {
		g.Logger.Errorf("failed to copy %d assets:\n- %s", len(allErrors), strings.Join(allErrors, "\n- "))
	}

	g.Logger.Printf("Asset copy took %s", time.Since(start))
}
