package generators

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/log"

	"gobsidian/internal/builders"
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

const (
	assetsDir           = "assets"
	generatedDir        = "generated"
	srcDir              = "src"
	imagesDir           = "images"
	fileTreeFilename    = "explorer.json"
	graphFilename       = "graph.json"
	searchIndexFilename = "search-index.json"
	tagsFilename        = "tags.json"
	chromaCSSPath       = "chroma.css"
	diagnosticsFile     = "vault-diagnostics.json"
)

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

	fileTreeBuilder := builders.NewFileTreeBuilder(g.Logger)
	notesRepository := repository.NewNoteRepository()
	graphBuilder := builders.NewGraphBuilder(g.Logger, g.SiteConfig.OutputDirectory, generatedDir, graphFilename, notesRepository)
	tagsBuilder := builders.NewTagsBuilder(g.Logger, g.SiteConfig.OutputDirectory, generatedDir, tagsFilename)
	searchIndexBuilder := builders.NewSearchIndexBuilder(g.Logger, g.SiteConfig.OutputDirectory, generatedDir, searchIndexFilename)
	transformCtx := transformers.NewTransformContext(g.SiteConfig.BaseURL, g.SiteConfig.OutputDirectory, notesRepository)

	pipeline := g.newTransformationPipeline()

	// Parse notes and get all notes by path
	parseTime := g.Parser.ParseNotes(scannedNotes, notesRepository)
	notesByPath := notesRepository.GetAllByPath()
	notesSlice := notesRepository.GetAllNotesSlice()
	fmt.Println(printer.Print("Parsed notes", "count", len(scannedNotes), "duration", parseTime))

	// Run transformations
	transformTime := g.runTransformations(pipeline, transformCtx, notesSlice)
	fmt.Println(printer.Print("Transformed content", "count", len(notesSlice), "duration", transformTime))

	// Build graph
	graphTime := graphBuilder.Build(notesByPath)
	fmt.Println(printer.Print("Built graph", "count", len(notesByPath), "duration", graphTime))
	if err := graphBuilder.ExportAndSaveAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save graph", "error", err))
		return err
	}
	fmt.Println(printer.Success("Exported graph"))

	// Build folder tree
	treeBuild, err := fileTreeBuilder.Build(notesByPath)
	if err != nil {
		fmt.Println(printer.Error("Failed to build folder tree", "error", err))
		return fmt.Errorf("failed to build folder tree: %w", err)
	}
	fmt.Println(printer.Print("Built folder tree", "count", treeBuild.NumberOfFolders, "duration", treeBuild.Duration))

	// Sort notes
	sortedNotes := g.sortNotes(notesSlice)
	templateExecutor := executors.NewTemplateExecutor(
		&g.SiteConfig,
		g.Logger,
		g.Templates,
	)

	// Generate breadcrumbs
	g.generateBreadcrumbs(sortedNotes)

	// Execute templates
	templateExecTime, numberOfTags, err := templateExecutor.Execute(sortedNotes, treeBuild.Root)
	if err != nil {
		fmt.Println(printer.Error("Failed to execute templates", "error", err))
		return err
	}
	fmt.Println(printer.Print("Executed templates", "count", len(sortedNotes), "duration", templateExecTime))

	// Build tags
	tagsTime := tagsBuilder.Build(sortedNotes)
	fmt.Println(printer.Print("Built tags", "count", "duration", tagsTime))
	if err := tagsBuilder.ExportAndSaveAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save tags", "error", err))
		return err
	}
	fmt.Println(printer.Success("Exported tags"))

	// Build search index
	searchIndexTime := searchIndexBuilder.Build(sortedNotes)
	fmt.Println(printer.Print("Built search index", "count", len(sortedNotes), "duration", searchIndexTime))
	if err := searchIndexBuilder.ExportAndSaveAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save search index", "error", err))
		return err
	}
	fmt.Println(printer.Success("Exported search index"))

	// Copy assets
	assetCopyTime := g.copyAssets(notesByPath, treeBuild.JSON)
	err = g.generateSyntaxHighlighterCSS()
	if err != nil {
		fmt.Println(printer.Error("Failed to generate syntax highlighter CSS", "error", err))
		return err
	}
	fmt.Println(printer.Print("Copied assets", "count", len(notesByPath), "duration", assetCopyTime))

	// Print diagnostics summary
	if g.shouldPrintVaultHealth {
		g.printDiagnosticsSummary(notesSlice, printer)
	}

	// Log summary
	g.Logger.Debug("Site generation complete!",
		"duration", time.Since(start),
		"notes", len(notesByPath),
		"notes/second", float64(len(notesByPath))/time.Since(start).Seconds(),
		"tags", numberOfTags,
		"preview pages", len(notesByPath),
		"folders", treeBuild.NumberOfFolders,
	)
	return nil
}

func (g *StaticSiteGenerator) newTransformationPipeline() *transformers.Pipeline {
	return transformers.NewPipeline(
		transformers.NewImageProcessor(g.Regexes.ObsidianImageRegex),
		transformers.NewHashtagEnricher(g.Regexes.HashtagRegex),
		transformers.NewCalloutTransformer(g.Regexes.ObsidianCalloutRegex),
		transformers.NewWikilinkTransformer(
			g.Regexes.WikilinkRegex,
			g.Regexes.ObsidianImageRegex,
			g.Regexes.HashtagRegex,
			g.Regexes.ObsidianCalloutRegex,
			g.Regexes.FootnoteRegex,
			g.Logger,
			g.MarkdownRenderer,
		),
		transformers.NewSyntaxHighlighter(true),
		transformers.NewFootnoteTransformer(g.Regexes.FootnoteRegex, g.Logger),
	)
}

func (g *StaticSiteGenerator) runTransformations(
	pipeline *transformers.Pipeline,
	ctx *transformers.TransformContext,
	notes []*models.ParsedNote,
) time.Duration {
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
					finalHTMLBytes = bytes.Replace(
						finalHTMLBytes,
						[]byte(embeddedPost.ID),
						[]byte(embeddedPost.Content),
						1,
					)
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

	return time.Since(start)
}

func (g *StaticSiteGenerator) generateBreadcrumbs(notes []*models.ParsedNote) {
	for _, note := range notes {
		breadcrumbs := []models.Breadcrumb{}
		parts := strings.Split(note.URL, "/")
		noteURL := ""

		for i, part := range parts {
			if part == "" {
				continue
			}

			noteURL += "/" + part

			part = strings.ReplaceAll(part, ".html", "")
			part = utils.Deslugify(part)

			isLast := i == len(parts)-1

			breadcrumbs = append(breadcrumbs, models.Breadcrumb{
				Title:  part,
				URL:    noteURL,
				IsLast: isLast,
			})
		}

		note.Breadcrumbs = breadcrumbs
	}
}

func (g *StaticSiteGenerator) computeTags(
	notes []*models.ParsedNote,
) ([]models.Tag, map[string][]*models.ParsedNote) {
	tagsMap := make(map[string]*models.Tag)
	for _, note := range notes {
		for _, tag := range note.Tags {
			tagsMap[tag.Slug] = &tag
			tagsMap[tag.Slug].Count++
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
