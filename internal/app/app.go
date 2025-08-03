package app

import (
	"flag"
	"fmt"
	"os"

	"gobsidian/internal/config"
	"gobsidian/internal/crawler"
	"gobsidian/internal/diff"
	"gobsidian/internal/generators"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/renderers"
	"gobsidian/internal/server"
	"gobsidian/internal/terminal"

	"github.com/gomarkdown/markdown/html"
)

type App struct {
	Printer   *terminal.PrettyPrinter
	Config    *config.Config
	Flags     *models.CmdArgs
	Crawler   *crawler.VaultCrawler
	Generator *generators.StaticSiteGenerator

	obsidianParser   *parsers.ObsidianParser
	excalidrawParser *parsers.ExcalidrawParser
	parserManager    *parsers.Manager
	renderer         *html.Renderer
}

func NewApp() *App {
	return &App{
		Printer: terminal.NewPrettyPrinter(),
	}
}

func (app *App) InitApp() error {
	if err := app.ParseFlags(); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	if err := app.LoadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if app.Flags.InputDirectory != "" {
		app.Config.SiteConfig.InputDirectory = app.Flags.InputDirectory
	}

	if app.Flags.OutputDirectory != "" {
		app.Config.SiteConfig.OutputDirectory = app.Flags.OutputDirectory
	}

	app.obsidianParser = parsers.NewObsidianParser(&parsers.ObsidianParser{
		InputDirectory: app.Config.SiteConfig.InputDirectory,
		Regexes:        &app.Config.Regexes,
	})
	app.excalidrawParser = parsers.NewExcalidrawParser(
		app.Config.SiteConfig.InputDirectory,
		app.Config.SiteConfig.OutputDirectory,
		app.Config.SiteConfig.BaseURL,
		app.Config.SiteConfig.Port,
		&app.Config.Regexes,
	)

	app.parserManager = parsers.NewManager(
		app.Config.SiteConfig,
		map[models.NoteType]parsers.Parser{
			models.NoteTypeMarkdown:   app.obsidianParser,
			models.NoteTypeExcalidraw: app.excalidrawParser,
		},
	)
	app.renderer = renderers.NewMarkdownRenderer()

	app.Crawler = crawler.NewVaultCrawler(app.Config.SiteConfig.InputDirectory, app.Config.SiteConfig.LandingPage, &app.Config.Regexes)

	gen, err := generators.NewStaticSiteGenerator(
		app.Config.SiteConfig,
		&app.Config.Regexes,
		app.Config.Logger,
		app.Config.Templates,
		app.Crawler,
		app.renderer,
		app.Flags.VaultHealth,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize generator: %w", err)
	}

	app.Generator = gen

	return nil
}

func (app *App) ParseFlags() error {
	flags := &models.CmdArgs{}
	flag.StringVar(&flags.InputDirectory, "vault", "", "Vault directory to scan")
	flag.StringVar(&flags.CPUProfile, "cpuprofile", "", "write cpu profile to `file`")
	flag.StringVar(&flags.MemProfile, "memprofile", "", "write memory profile to `file`")
	flag.BoolVar(&flags.Clear, "clear", false, "Clear the public folder before generating")
	flag.BoolVar(&flags.Serve, "serve", false, "Serve the generated website after building")
	flag.StringVar(&flags.Port, "port", "8000", "Port to serve the website on")
	flag.BoolVar(&flags.VaultHealth, "vhealth", false, "Check the health of the vault")
	flag.Parse()

	app.Flags = flags
	return nil
}

func (app *App) LoadConfig() error {
	cfg, err := config.LoadConfig(".config.toml")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	app.Config = &cfg
	return nil
}

func (app *App) ClearOutputDirectory() error {
	if err := os.RemoveAll(app.Config.SiteConfig.OutputDirectory); err != nil {
		return fmt.Errorf("failed to clear output directory: %w", err)
	}
	fmt.Println(app.Printer.Info("Output directory cleared"))
	return nil
}

func (app *App) DetectChanges() ([]string, error) {
	diffTracker := diff.NewTracker(
		app.Config.SiteConfig.InputDirectory,
		app.Config.SiteConfig.OutputDirectory,
	)

	hasChanges, err := diffTracker.HasChanges()
	if err != nil {
		return nil, fmt.Errorf("failed to check for changes: %w", err)
	}

	if !hasChanges {
		return nil, nil
	}

	changedFiles, err := diffTracker.GetChangedFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to get changed files: %w", err)
	}

	_, err = diffTracker.UpdateLockFile()
	if err != nil {
		return nil, fmt.Errorf("failed to update lock file: %w", err)
	}

	// fmt.Println(app.Printer.PrintDiffsSummary(diffs))
	// fmt.Println(app.Printer.PrintDiffsDetailed(diffs))
	return changedFiles, nil
}

func (app *App) ServeSite() {
	webServer := server.NewWebServer(app.Flags.Port, app.Config.SiteConfig.OutputDirectory, app.Config.Logger)
	fmt.Println(app.Printer.Info("Serving website", "port", app.Flags.Port))
	webServer.Serve()
}
