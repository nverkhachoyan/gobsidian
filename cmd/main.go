package main

import (
	"fmt"
	"os"
	"time"

	"gobsidian/internal/app"
	"gobsidian/internal/profilers"
	"gobsidian/internal/terminal"
	"gobsidian/internal/utils"
)

func main() {
	startAppTime := time.Now()

	app := app.NewApp()
	if err := app.InitApp(); err != nil {
		fmt.Println(app.Printer.Error("Failed to initialize app", "error", err))
		os.Exit(1)
	}

	reporter := terminal.NewProgressReporter(app.Printer)
	profilers.RunProfilers(&app.Flags.CPUProfile, &app.Flags.MemProfile)

	if app.Flags.Clear {
		clearStart := time.Now()
		if err := app.ClearOutputDirectory(); err != nil {
			fmt.Println(app.Printer.Error("Failed to clear output directory", "error", err))
			os.Exit(1)
		}
		clearStep := terminal.Step{Name: "Cleared output directory", Icon: "🧹"}
		reporter.ReportStep(clearStep, terminal.StepResult{
			Duration: time.Since(clearStart),
			Count:    0,
			Error:    nil,
		})
	}

	isOutputDirectoryExists, err := utils.IsDirectoryExists(app.Config.SiteConfig.OutputDirectory)
	if err != nil {
		fmt.Println(app.Printer.Error("Failed to check if output directory exists", "error", err))
		os.Exit(1)
	}

	if !isOutputDirectoryExists {
		err := os.Mkdir(app.Config.SiteConfig.OutputDirectory, 0755)
		if err != nil {
			fmt.Println(app.Printer.Error("Failed to create output directory", "error", err))
			os.Exit(1)
		}
	}

	changedFiles, err := app.DetectChanges()
	if err != nil {
		fmt.Println(app.Printer.Error("Failed to detect changes", "error", err))
		os.Exit(1)
	}

	scanStart := time.Now()
	var scanDuration time.Duration

	if !isOutputDirectoryExists {
		fmt.Println(app.Printer.Info("Output directory does not exist, scanning all notes"))
		scanDuration, err = app.Scanner.ScanAllNotes()
	} else if len(changedFiles) > 0 {
		fmt.Println(app.Printer.Info("Changes detected, scanning notes"))
		scanDuration, err = app.Scanner.ScanNotesByPaths(changedFiles)
	} else {
		fmt.Println(app.Printer.Info("No changes detected"))
	}

	scanStep := terminal.Step{Name: "Scanned files", Icon: "📄"}
	if err != nil {
		reporter.ReportStep(scanStep, terminal.StepResult{
			Duration: time.Since(scanStart),
			Count:    0,
			Error:    err,
		})
		fmt.Println(app.Printer.Error("Failed to scan notes", "error", err))
		os.Exit(1)
	}

	scannedNotes := app.Scanner.GetAllNotes()

	if len(scannedNotes) > 0 {
		reporter.ReportStep(scanStep, terminal.StepResult{
			Duration: scanDuration,
			Count:    len(scannedNotes),
			Error:    nil,
		})
		if err := app.Generator.Generate(scannedNotes); err != nil {
			fmt.Println(app.Printer.Error("Failed to run generator", "error", err))
			os.Exit(1)
		}
	}

	timeToGenerate := time.Since(startAppTime)
	fmt.Println(app.Printer.Success("Site generated successfully", "duration", timeToGenerate))

	if app.Flags.Serve {
		app.ServeSite()
	}

	os.Exit(0)
}
