package generators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/log"

	"gobsidian/internal/builders"
	"gobsidian/internal/config"
	"gobsidian/internal/crawler"
	"gobsidian/internal/executors"
	"gobsidian/internal/models"
	"gobsidian/internal/terminal"
	"gobsidian/internal/transformers"
	"gobsidian/internal/utils"

	"github.com/gomarkdown/markdown"

	"github.com/gomarkdown/markdown/html"
)

type StaticSiteGenerator struct {
	SiteConfig             config.SiteConfig
	Regexes                *config.Regexes
	Logger                 *log.Logger
	Templates              *template.Template
	Crawler                *crawler.VaultCrawler
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
	excalidrawsDir      = "excalidraws"
)

type ExcalidrawBase64 struct {
	Data string `json:"data"`
}

func NewStaticSiteGenerator(
	siteConfig config.SiteConfig,
	regexes *config.Regexes,
	logger *log.Logger,
	templates *template.Template,
	crawler *crawler.VaultCrawler,
	markdownRenderer *html.Renderer,
	shouldPrintVaultHealth bool,
) (*StaticSiteGenerator, error) {
	return &StaticSiteGenerator{
		SiteConfig:             siteConfig,
		Regexes:                regexes,
		Logger:                 logger,
		Templates:              templates,
		Crawler:                crawler,
		MarkdownRenderer:       markdownRenderer,
		shouldPrintVaultHealth: shouldPrintVaultHealth,
	}, nil
}

func (g *StaticSiteGenerator) Generate() error {
	start := time.Now()
	printer := terminal.NewPrettyPrinter()
	explorerBuilder := builders.NewExplorerBuilder(g.Logger)
	graphBuilder := builders.NewGraphBuilder(g.Logger, g.SiteConfig.OutputDirectory, generatedDir, graphFilename)
	tagsBuilder := builders.NewTagsBuilder(g.Logger, g.SiteConfig.OutputDirectory, generatedDir, tagsFilename)
	searchIndexBuilder := builders.NewSearchIndexBuilder(g.Logger, g.SiteConfig.OutputDirectory, generatedDir, searchIndexFilename)

	// Crawl vault
	rootNode, err := g.Crawler.Crawl()
	if err != nil {
		fmt.Println(printer.Error("Failed to crawl vault", err))
	}

	fileIndex := g.Crawler.GetFileIndex()
	tagIndex := g.Crawler.GetTagIndex()

	// Run Transformations
	transformTime := g.runTransformations(rootNode)
	fmt.Println(printer.Success("Transformed files.", "duration", transformTime))

	// Build Graph
	graphBuildTime := graphBuilder.Build(fileIndex)
	fmt.Println(printer.Print("Built network graph", "duration", graphBuildTime))
	if err := graphBuilder.ExportAndSaveAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save graph", "error", err))
		return err
	}
	if err := graphBuilder.ExportAndSaveLocalGraphsAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save local graphs", "error", err))
		return err
	}
	fmt.Println(printer.Success("Exported local graphs"))

	// Build tags
	tagsBuildTime := tagsBuilder.Build(fileIndex)
	if err := tagsBuilder.ExportAndSaveAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save tags", "error", err))
		return err
	}
	fmt.Println(printer.Print("Built tags", "duration", tagsBuildTime))

	// Build Search Index
	searchIndexTime := searchIndexBuilder.Build(fileIndex)
	if err := searchIndexBuilder.ExportAndSaveAsJSON(); err != nil {
		fmt.Println(printer.Error("Failed to export and save tags", "error", err))
		return err
	}
	fmt.Println(printer.Print("Built search index", "duration", searchIndexTime))

	explorerResult := explorerBuilder.Build(rootNode)
	g.exportFiletreeJSON(explorerResult.JSON)
	fmt.Println(printer.Print("Built file tree", "duration", explorerResult.Duration))

	err = g.generateSyntaxHighlighterCSS()
	if err != nil {
		fmt.Println(printer.Error("Failed to generate syntax highlighter CSS", "error", err))
		return err
	}

	// ASSETS
	// TODO: fix issue where images don't exist and we try to copy
	imagesCopyTime := g.copyImages(fileIndex)
	fmt.Println(printer.Print("Copied assets", "duration", imagesCopyTime))

	assetsDirCopyTime := g.copyAssetsDir()
	fmt.Println(printer.Print("Copied assets", "duration", assetsDirCopyTime))

	// Generate breadcrumbs
	g.generateBreadcrumbs(fileIndex)
	fmt.Println(printer.Print("Generated breadcrumbs"))

	// // Copy excalidraws
	excalidrawCopyTime := g.createExcalidrawJSONs(fileIndex)
	fmt.Println(printer.Print("Copied excalidraws", "duration", excalidrawCopyTime))

	// TODO: impl
	// Print diagnostics summary
	// if g.shouldPrintVaultHealth {
	// 	g.printDiagnosticsSummary(notesSlice, printer)
	// }

	// Execute templates
	templateExecutor := executors.NewTemplateExecutor(rootNode, fileIndex, tagIndex, &g.SiteConfig, g.Logger, g.Templates)
	templateExecTime, _, err := templateExecutor.Execute()
	if err != nil {
		fmt.Println(printer.Error("Failed to execute templates", "error", err))
		return err
	}
	fmt.Println(printer.Print("Executed templates", "duration", templateExecTime))

	// Log summary
	g.Logger.Debug("Site generation complete!",
		"duration", time.Since(start),
		"notes", len(fileIndex),
		"notes/second", float64(len(fileIndex))/time.Since(start).Seconds(),
	)
	return nil
}

func (g *StaticSiteGenerator) runTransformations(rootNode *models.VaultNode) time.Duration {
	start := time.Now()
	ctx := transformers.NewTransformContext(g.SiteConfig.BaseURL, g.SiteConfig.Port, g.SiteConfig.OutputDirectory)
	pipeline := transformers.NewPipeline(
		transformers.NewExcalidrawEnricher(),
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
		transformers.NewFootnoteEmbeddedTransformer(g.Regexes.FootnoteRegex, g.Logger),
		transformers.NewSyntaxHighlighter(true),
	)
	g.transformNode(pipeline, ctx, rootNode)
	return time.Since(start)
}

func (g *StaticSiteGenerator) transformNode(
	pipeline *transformers.Pipeline,
	ctx *transformers.TransformContext,
	node *models.VaultNode,
) {
	if node.GetNoteType() == models.NoteTypeMarkdown || node.GetNoteType() == models.NoteTypeExcalidraw {
		nodeCtx := ctx.Clone()
		transformedContent, err := pipeline.Transform(node.Markdown, node, nodeCtx)
		if err != nil {
			g.Logger.Warn("failure to transform", "error", err)
			return
		}
		finalHTMLBytes := markdown.ToHTML([]byte(transformedContent), nil, g.MarkdownRenderer)

		for _, embeddedPost := range nodeCtx.EmbeddedPosts {
			finalHTMLBytes = bytes.Replace(
				finalHTMLBytes,
				[]byte(embeddedPost.ID),
				[]byte(embeddedPost.Content),
				1,
			)
		}

		node.HTML = template.HTML(finalHTMLBytes)

	}

	for _, child := range node.Children {
		g.transformNode(pipeline, ctx, child)
	}
}

func (g *StaticSiteGenerator) generateBreadcrumbs(fileIndex map[string]*models.VaultNode) {
	for _, node := range fileIndex {
		breadcrumbs := []models.Breadcrumb{}
		parts := strings.Split(node.URL, string(filepath.Separator))
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

		node.Breadcrumbs = breadcrumbs
	}
}

func (g *StaticSiteGenerator) createExcalidrawJSONs(fileIndex map[string]*models.VaultNode) time.Duration {
	start := time.Now()

	for _, note := range fileIndex {
		if note.NoteType != models.NoteTypeExcalidraw {
			continue
		}

		slugifiedFileName := utils.Slugify(note.BaseName)
		excalidrawPath := filepath.Join(g.SiteConfig.OutputDirectory, generatedDir, excalidrawsDir, slugifiedFileName+".json")

		excalidrawJSONBytes, err := json.Marshal(note.ExcalidrawPayload)
		if err != nil {
			fmt.Printf("failed to marshal excalidraw JSON: %v\n", err)
			g.Logger.Error("failed to marshal excalidraw JSON", "error", err)
			return time.Since(start)
		}
		if err := os.WriteFile(excalidrawPath, excalidrawJSONBytes, 0644); err != nil {
			fmt.Printf("failed to write excalidraw JSON: %v\n", err)
			g.Logger.Error("failed to write excalidraw JSON", "error", err)
			return time.Since(start)
		}
	}

	return time.Since(start)
}
