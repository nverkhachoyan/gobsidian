package generators

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	"sync"
	"time"

	"github.com/charmbracelet/log"

	"os"
	"path/filepath"
	"sort"
	"strings"

	"gobsidian/internal/executors"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/utils"

	"github.com/google/uuid"

	"github.com/gomarkdown/markdown/html"

	"github.com/gomarkdown/markdown"
)

type StaticSiteGenerator struct {
	InputDirectory      string
	OutputDirectory     string
	SiteTitle           string
	SiteSubtitle        string
	BaseURL             string
	Regexes             *models.ParserRegexes
	Logger              *log.Logger
	Templates           *template.Template
	Parser              parsers.Parser
	MarkdownRenderer    *html.Renderer
	RenderedPostCache   map[string]template.HTML
	RenderingInProgress map[string]bool
	cacheMu             sync.RWMutex
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
		InputDirectory:      cfg.InputDirectory,
		OutputDirectory:     cfg.OutputDirectory,
		SiteTitle:           cfg.SiteTitle,
		SiteSubtitle:        cfg.SiteSubtitle,
		BaseURL:             cfg.BaseURL,
		Regexes:             cfg.Regexes,
		Logger:              cfg.Logger,
		Templates:           cfg.Templates,
		Parser:              cfg.Parser,
		MarkdownRenderer:    cfg.MarkdownRenderer,
		RenderedPostCache:   make(map[string]template.HTML),
		RenderingInProgress: make(map[string]bool),
	}, nil
}

func (g *StaticSiteGenerator) GenerateSite() error {
	start := time.Now()
	g.Logger.Print("Generating site...")

	discoveryStart := time.Now()
	noteScanner, err := g.walkAndScanNotes()
	if err != nil {
		return fmt.Errorf("failed to build note registry %w", err)
	}
	g.Logger.Printf("Note discovery took %s", time.Since(discoveryStart))

	if err := g.createOutputDirectories(); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	wikilinkStart := time.Now()
	wikilinkResolver, err := g.parseNotes(noteScanner)
	if err != nil {
		return fmt.Errorf("error during initial walk: %w", err)
	}
	g.Logger.Printf("Wikilink parsing took %s", time.Since(wikilinkStart))

	notesByPath := wikilinkResolver.GetAllByPath()

	transformationStart := time.Now()
	g.applyTransformationsConcurrently(notesByPath, wikilinkResolver)
	g.Logger.Printf("Transformation took %s", time.Since(transformationStart))

	buildDirs := time.Now()
	fileTree, numberOfFolders, err := g.buildFolderTree(notesByPath)
	if err != nil {
		return fmt.Errorf("failed to create subdirectories: %w", err)
	}
	g.Logger.Printf("Directory creation took %s", time.Since(buildDirs))

	graphStart := time.Now()
	graphGenerator := NewGraphGenerator(&GraphGenerator{
		Logger: g.Logger,
	})
	graph := graphGenerator.Generate(notesByPath, wikilinkResolver)
	g.Logger.Printf("Graph generation took %s", time.Since(graphStart))

	notes := make([]*models.ParsedNote, 0, len(notesByPath))
	for _, note := range notesByPath {
		notes = append(notes, note)
	}

	templateStart := time.Now()
	numberOfTags, err := g.executeTemplatesBatch(notes, fileTree, graph)
	if err != nil {
		return err
	}
	g.Logger.Printf("Template execution took %s", time.Since(templateStart))

	assetCopyStart := time.Now()
	g.copyAssets(notesByPath)
	g.Logger.Printf("Asset copy took %s", time.Since(assetCopyStart))

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

func (g *StaticSiteGenerator) walkAndScanNotes() (
	*NoteScanner,
	error,
) {
	registry := NewNoteScanner(g.InputDirectory)
	err := filepath.WalkDir(g.InputDirectory, func(path string, info os.DirEntry, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			notePath := strings.TrimPrefix(path, g.InputDirectory+"/")
			isInsideFolder := strings.Contains(notePath, "/")
			note := &models.ShallowNote{
				FileName:       info.Name(),
				FullPath:       path,
				RelativePath:   notePath,
				IsInsideFolder: isInsideFolder,
			}
			registry.AddNote(note)

		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error during initial walk: %w", err)
	}

	return registry, nil
}

func (g *StaticSiteGenerator) applyTransformationsConcurrently(notesByPath map[string]*models.ParsedNote, wikilinkResolver *WikilinkResolver) error {
	notes := make([]*models.ParsedNote, 0, len(notesByPath))
	for _, note := range notesByPath {
		notes = append(notes, note)
	}

	numWorkers := min(runtime.NumCPU(), len(notes))
	notesChan := make(chan *models.ParsedNote, len(notes))
	errorsChan := make(chan error, len(notes))

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for note := range notesChan {
				if err := g.applyTransformations(note, wikilinkResolver); err != nil {
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

	return nil
}

func (g *StaticSiteGenerator) applyTransformations(note *models.ParsedNote, wikilinkResolver *WikilinkResolver) error {
	postBody := string(note.RawBody)
	contentWithImagesReplaced := g.replaceObsidianImageLinks(postBody)

	contentWithWikilinks, embeddedPosts := g.replaceWikilinksAndResolveBacklinks(
		contentWithImagesReplaced,
		note,
		wikilinkResolver,
	)
	hashtagAsLinks := g.enrichHashtagsWithLinks(contentWithWikilinks)

	finalHTMLBytes := markdown.ToHTML([]byte(hashtagAsLinks), nil, g.MarkdownRenderer)

	for _, embeddedPost := range embeddedPosts {
		finalHTMLBytes = bytes.Replace(finalHTMLBytes, []byte(embeddedPost.ID), []byte(embeddedPost.Content), 1)
	}

	note.HTMLContent = template.HTML(finalHTMLBytes)
	return nil
}

func (g *StaticSiteGenerator) replaceObsidianImageLinks(body string) string {
	return g.Regexes.ObsidianImageRegex.ReplaceAllStringFunc(string(body), func(match string) string {
		parts := g.Regexes.ObsidianImageRegex.FindStringSubmatch(match)
		if len(parts) > 3 {
			imageName := parts[1]
			width := parts[2]
			height := parts[3]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s" width="%s" height="%s" />`, imageName, imageName, width, height)
		}
		if len(parts) > 2 {
			imageName := parts[1]
			width := parts[2]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s"  width="%s" />`, imageName, imageName, width)
		}
		if len(parts) > 1 {
			imageName := parts[1]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s"  />`, imageName, imageName)
		}

		return match
	})
}

func (g *StaticSiteGenerator) buildFolderTree(notes map[string]*models.ParsedNote) (*models.Folder, int, error) {
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
	return root, numberOfFolders, nil
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

func (g *StaticSiteGenerator) renderPostContent(
	post *models.ParsedNote,
	wikilinkResolver *WikilinkResolver,
) template.HTML {
	if renderedPost, ok := g.getCachedPost(post.URL); ok {
		return renderedPost
	}

	if g.isRenderingInProgress(post.URL) {
		g.Logger.Warn("Circular embed detected, returning placeholder/error for", "title", post.Title)
		return template.HTML(fmt.Sprintf(`<div class="embedded-post-circular-error">Circular Embed: %s</div>`, post.Title))
	}

	g.setRenderingInProgress(post.URL, true)

	contentWithImagesReplaced := g.replaceObsidianImageLinks(string((*post).RawBody))
	contentWithWikilinks, embeddedPosts := g.replaceWikilinksAndResolveBacklinks(
		contentWithImagesReplaced,
		post,
		wikilinkResolver,
	)
	hashtagAsLinks := g.enrichHashtagsWithLinks(contentWithWikilinks)

	finalHTMLBytes := markdown.ToHTML([]byte(hashtagAsLinks), nil, g.MarkdownRenderer)

	for _, embeddedPost := range embeddedPosts {
		finalHTMLBytes = bytes.Replace(finalHTMLBytes, []byte(embeddedPost.ID), []byte(embeddedPost.Content), 1)
	}

	g.setCachedPost(post.URL, template.HTML(finalHTMLBytes))
	g.setRenderingInProgress(post.URL, false)

	return template.HTML(finalHTMLBytes)
}

func (g *StaticSiteGenerator) replaceWikilinksAndResolveBacklinks(
	body string,
	pd *models.ParsedNote,
	wikilinkResolver *WikilinkResolver,
) (string, map[string]models.EmbeddedPost) {
	embeddedPosts := make(map[string]models.EmbeddedPost)

	return g.Regexes.WikilinkRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := g.Regexes.WikilinkRegex.FindStringSubmatch(match)
		matchedLinkText := parts[1]
		pseudoName := matchedLinkText

		if len(parts) > 2 && parts[2] != "" {
			pseudoName = parts[2]
		}

		isEmbedded := strings.HasPrefix(match, "!")

		if targetNote, found := wikilinkResolver.ResolveWikilink(matchedLinkText); found {
			wikilinkResolver.AddBacklink(targetNote, pd, pseudoName)

			if isEmbedded {
				return g.handleEmbeddedPost(wikilinkResolver, embeddedPosts, targetNote)
			}
			return fmt.Sprintf(`<a href="%s">%s</a>`, targetNote.URL, pseudoName)
		}

		return fmt.Sprintf(`<a href="/%s">%s</a>`, utils.Slugify(matchedLinkText)+".html", pseudoName)
	}), embeddedPosts
}

func (g *StaticSiteGenerator) handleEmbeddedPost(
	wikilinkResolver *WikilinkResolver,
	embeddedPosts map[string]models.EmbeddedPost,
	embeddedPost *models.ParsedNote,
) string {
	var uniqueID string

	htmlContent := g.renderPostContent(embeddedPost, wikilinkResolver)
	uniqueID = uuid.New().String()
	embeddedPosts[uniqueID] = models.EmbeddedPost{
		ID:      uniqueID,
		Content: htmlContent,
	}
	return fmt.Sprintf(`<div class="embedded-post">
		<div class="embedded-post-header">
		<a href="%s"><h2>%s</h2></a>
		</div>
		%s</div>`, embeddedPost.URL, embeddedPost.Title, uniqueID)
}

func (g *StaticSiteGenerator) enrichHashtagsWithLinks(body string) string {
	return g.Regexes.HashtagRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := g.Regexes.HashtagRegex.FindStringSubmatch(match)
		tagName := parts[1]
		return fmt.Sprintf(`<a class="tag" href="/tag/%s">#%s</a>`, tagName, tagName)
	})
}

func (g *StaticSiteGenerator) executeTemplatesBatch(notes []*models.ParsedNote, fileTree *models.Folder, graph []byte) (int, error) {
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

	g.Logger.Debug("Finished writing template files")
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

func (g *StaticSiteGenerator) createOutputDirectories() error {
	if err := os.MkdirAll(filepath.Join(g.OutputDirectory, "images"), 0755); err != nil {
		return fmt.Errorf("failed to create images directory %s: %w", filepath.Join(g.OutputDirectory, "images"), err)
	}
	return nil
}

type parseResult struct {
	note *models.ParsedNote
	err  error
}

func (g *StaticSiteGenerator) parseNotes(noteScanner *NoteScanner) (
	*WikilinkResolver,
	error,
) {
	shallowNotes := noteScanner.GetAllNotes()

	numWorkers := min(runtime.NumCPU(), len(shallowNotes))
	notesChan := make(chan *models.ShallowNote, len(shallowNotes))
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

	wikilinkResolver := NewWikilinkResolver()

	for result := range resultsChan {
		if result.err != nil {
			g.Logger.Warn("failed to parse note", "error", result.err)
			continue
		}

		wikilinkResolver.AddNote(result.note)
	}

	return wikilinkResolver, nil
}

func (g *StaticSiteGenerator) parseNoteWorker(shallowNote *models.ShallowNote) parseResult {
	if shallowNote == nil {
		return parseResult{
			note: nil,
			err:  fmt.Errorf("received nil shallowNote"),
		}
	}

	parsedNoteFromParser, err := g.Parser.ParseNote(shallowNote.FullPath, shallowNote.FileName)

	// passing as pointer is safe since each parseNoteWorker gets a different note
	parsedNote := &models.ParsedNote{
		ShallowNote: shallowNote,
	}

	if err != nil {
		g.Logger.Warn("failed to parse note", "path", shallowNote.FullPath, "error", err)
		return parseResult{
			note: parsedNote,
			err:  err,
		}
	}

	URL := strings.TrimPrefix(shallowNote.FullPath, g.InputDirectory)
	URL = strings.TrimPrefix(URL, "/")   // Remove leading slash
	URL = strings.TrimSuffix(URL, ".md") // Remove .md extension

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
}

func (g *StaticSiteGenerator) getCachedPost(url string) (template.HTML, bool) {
	g.cacheMu.RLock()
	defer g.cacheMu.RUnlock()
	post, ok := g.RenderedPostCache[url]
	return post, ok
}

func (g *StaticSiteGenerator) setCachedPost(url string, content template.HTML) {
	g.cacheMu.Lock()
	defer g.cacheMu.Unlock()
	g.RenderedPostCache[url] = content
}

func (g *StaticSiteGenerator) isRenderingInProgress(url string) bool {
	g.cacheMu.RLock()
	defer g.cacheMu.RUnlock()
	return g.RenderingInProgress[url]
}

func (g *StaticSiteGenerator) setRenderingInProgress(url string, inProgress bool) {
	g.cacheMu.Lock()
	defer g.cacheMu.Unlock()
	g.RenderingInProgress[url] = inProgress
}
