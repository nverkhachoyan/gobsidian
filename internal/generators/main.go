package generators

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
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
	assetsDir        = "assets"
	cssDir           = "css"
	srcDir           = "src"
	imagesDir        = "images"
	fileTreeFilename = "file-tree.json"
	chromaCSSPath    = "assets/chroma.css"
	diagnosticsFile  = "vault-diagnostics.json"
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
	reporter := terminal.NewProgressReporter(printer)
	fileTreeBuilder := builders.NewFileTreeBuilder(g.Logger)
	notesRepository := repository.NewNoteRepository()
	graphBuilder := builders.NewGraphBuilder(g.Logger)
	transformCtx := transformers.NewTransformContext(g.SiteConfig.BaseURL, g.SiteConfig.OutputDirectory, notesRepository)

	pipeline := g.newTransformationPipeline()

	// Start of the pipeline

	parseTime := g.Parser.ParseNotes(scannedNotes, notesRepository)
	notesByPath := notesRepository.GetAllByPath()
	notesSlice := notesRepository.GetAllNotesSlice()
	reporter.ReportStep(
		terminal.Step{Name: "Parsing notes", Icon: "📝"},
		terminal.StepResult{
			Duration: parseTime,
			Count:    len(notesSlice),
			Error:    nil,
		})

	transformTime := g.runTransformations(pipeline, transformCtx, notesSlice)
	reporter.ReportStep(
		terminal.Step{Name: "Transformed content", Icon: "✨"},
		terminal.StepResult{
			Duration: transformTime,
			Count:    len(notesSlice),
			Error:    nil,
		})

	graphTime, graphJSON := graphBuilder.Build(notesByPath, notesRepository)
	reporter.ReportStep(
		terminal.Step{Name: "Built graph", Icon: "🔗"},
		terminal.StepResult{
			Duration: graphTime,
			Count:    len(notesByPath),
			Error:    nil,
		})

	treeBuild, err := fileTreeBuilder.Build(notesByPath)
	if err != nil {
		reporter.ReportStep(
			terminal.Step{Name: "Failed to build folder tree", Icon: "📁"},
			terminal.StepResult{
				Duration: treeBuild.Duration,
				Count:    treeBuild.NumberOfFolders,
				Error:    err,
			})
		return fmt.Errorf("failed to build folder tree: %w", err)
	}
	reporter.ReportStep(
		terminal.Step{Name: "Built folder tree", Icon: "📁"},
		terminal.StepResult{
			Duration: treeBuild.Duration,
			Count:    treeBuild.NumberOfFolders,
			Error:    nil,
		})

	templateStart := time.Now()
	sortedNotes := g.sortNotes(notesSlice)
	templateExecutor := executors.NewTemplateExecutor(
		&g.SiteConfig,
		g.Logger,
		g.Templates,
	)

	templateExecTime, numberOfTags, err := templateExecutor.Execute(sortedNotes, treeBuild.Root, graphJSON, treeBuild.JSON)
	if err != nil {
		reporter.ReportStep(
			terminal.Step{Name: "Failed to execute templates", Icon: "🎨"},
			terminal.StepResult{
				Duration: time.Since(templateStart),
				Count:    0,
				Error:    err,
			})
		return err
	}
	reporter.ReportStep(
		terminal.Step{Name: "Executed templates", Icon: "🎨"},
		terminal.StepResult{
			Duration: templateExecTime,
			Count:    len(sortedNotes),
			Error:    nil,
		})

	assetCopyTime := g.copyAssets(notesByPath, treeBuild.JSON)
	err = g.generateSyntaxHighlighterCSS()
	if err != nil {
		reporter.ReportStep(
			terminal.Step{Name: "Failed to generate syntax highlighter CSS", Icon: "🎨"},
			terminal.StepResult{
				Duration: time.Since(templateStart),
				Count:    0,
				Error:    err,
			})
		return err
	}

	reporter.ReportStep(
		terminal.Step{Name: "Copied assets", Icon: "📦"},
		terminal.StepResult{
			Duration: assetCopyTime,
			Count:    len(notesByPath),
			Error:    nil,
		})

	if g.shouldPrintVaultHealth {
		g.printDiagnosticsSummary(notesSlice, printer)
	}

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
