package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/diff"
	"gobsidian/internal/generators"
	"gobsidian/internal/parsers"
	"gobsidian/internal/renderers"
	"gobsidian/internal/scanner"
	"gobsidian/internal/server"
	"gobsidian/internal/terminal"
)

type AppConfig struct {
	CPUProfile  string
	MemProfile  string
	Clear       bool
	Serve       bool
	Port        string
	VaultHealth bool
}

type App struct {
	printer *terminal.PrettyPrinter
	config  *config.Config
	flags   *AppConfig
}

func main() {
	app := &App{
		printer: terminal.NewPrettyPrinter(),
	}

	if err := app.run(); err != nil {
		fmt.Println(app.printer.Error("Application failed", "error", err))
		os.Exit(1)
	}
}

func (app *App) run() error {
	if err := app.parseFlags(); err != nil {
		return err
	}

	if err := app.loadConfig(); err != nil {
		return err
	}

	f, err := os.Create(app.flags.CPUProfile)
	if err != nil {
		app.config.Logger.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		app.config.Logger.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	fmem, err := os.Create(app.flags.MemProfile)
	if err != nil {
		app.config.Logger.Fatal("could not create memory profile: ", err)
	}
	defer fmem.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(fmem); err != nil {
		app.config.Logger.Fatal("could not write memory profile: ", err)
	}

	if app.flags.Clear {
		if err := app.clearOutputDirectory(); err != nil {
			return err
		}
	}

	changedFiles, err := app.detectChanges()
	if err != nil {
		return err
	}

	if err := app.generateSite(changedFiles); err != nil {
		return err
	}

	if app.flags.Serve {
		return app.serveSite()
	}

	return nil
}

func (app *App) parseFlags() error {
	flags := &AppConfig{}
	flag.StringVar(&flags.CPUProfile, "cpuprofile", "cpu.prof", "write cpu profile to `file`")
	flag.StringVar(&flags.MemProfile, "memprofile", "mem.prof", "write memory profile to `file`")
	flag.BoolVar(&flags.Clear, "clear", false, "Clear the public folder before generating")
	flag.BoolVar(&flags.Serve, "serve", false, "Serve the generated website after building")
	flag.StringVar(&flags.Port, "port", "8000", "Port to serve the website on")
	flag.BoolVar(&flags.VaultHealth, "vhealth", false, "Check the health of the vault")
	flag.Parse()

	app.flags = flags
	return nil
}

func (app *App) loadConfig() error {
	cfg, err := config.LoadConfig(".config.toml")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	app.config = &cfg
	return nil
}

func (app *App) clearOutputDirectory() error {
	if err := os.RemoveAll(app.config.SiteConfig.OutputDirectory); err != nil {
		return fmt.Errorf("failed to clear output directory: %w", err)
	}
	fmt.Println(app.printer.Info("Output directory cleared"))
	return nil
}

func (app *App) detectChanges() ([]string, error) {
	diffTracker := diff.NewTracker(app.config.SiteConfig.InputDirectory)

	hasChanges, err := diffTracker.HasChanges()
	if err != nil {
		return nil, fmt.Errorf("failed to check for changes: %w", err)
	}

	if !hasChanges {
		fmt.Println(app.printer.Info("No changes detected"))
		return nil, nil
	}

	fmt.Println(app.printer.Info("Changes detected, scanning notes"))

	changedFiles, err := diffTracker.GetChangedFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to get changed files: %w", err)
	}

	diffs, err := diffTracker.UpdateLockFile()
	if err != nil {
		return nil, fmt.Errorf("failed to update lock file: %w", err)
	}

	fmt.Println(app.printer.PrintDiffsSummary(diffs))
	return changedFiles, nil
}

func (app *App) generateSite(changedFiles []string) error {
	start := time.Now()

	reporter := terminal.NewProgressReporter(app.printer)

	scanner := scanner.NewNoteScanner(app.config.Logger, app.config.SiteConfig.InputDirectory)
	markdownRenderer := renderers.NewMarkdownRenderer()
	parser := parsers.NewObsidianParser(&parsers.ObsidianParser{
		InputDirectory: app.config.SiteConfig.InputDirectory,
		Regexes:        &app.config.Regexes,
	})

	scanStep := terminal.Step{Name: "Scanning files", Icon: "📄"}
	scanStart := time.Now()

	var scanDuration time.Duration
	var err error

	if len(changedFiles) > 0 {
		scanDuration, err = scanner.ScanNotesByPaths(changedFiles)
	} else {
		scanDuration, err = scanner.ScanAllNotes()
	}

	if err != nil {
		reporter.ReportStep(scanStep, terminal.StepResult{
			Duration: time.Since(scanStart),
			Count:    0,
			Error:    err,
		})
		return fmt.Errorf("failed to scan notes: %w", err)
	}

	scannedNotes := scanner.GetAllNotes()

	reporter.ReportStep(scanStep, terminal.StepResult{
		Duration: scanDuration,
		Count:    len(scannedNotes),
		Error:    nil,
	})

	gen, err := generators.NewStaticSiteGenerator(
		app.config.SiteConfig,
		&app.config.Regexes,
		app.config.Logger,
		app.config.Templates,
		parser,
		scanner,
		markdownRenderer,
		app.flags.VaultHealth,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize generator: %w", err)
	}

	if err := gen.Generate(scannedNotes); err != nil {
		return fmt.Errorf("failed to generate site: %w", err)
	}

	totalDuration := time.Since(start)
	fmt.Println(app.printer.Success("Site generated successfully", "duration", totalDuration))
	return nil
}

func (app *App) serveSite() error {
	webServer := server.NewWebServer(app.flags.Port, app.config.SiteConfig.OutputDirectory, app.config.Logger)
	fmt.Println(app.printer.Info("Serving website", "port", app.flags.Port))
	webServer.Serve()
	return nil
}
