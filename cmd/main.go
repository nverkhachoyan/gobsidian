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

	if err := app.Generator.Generate(); err != nil {
		fmt.Println(app.Printer.Error("Failed to run generator", "error", err))
		os.Exit(1)
	}

	if app.Flags.Serve {
		app.ServeSite()
	}

	os.Exit(0)
}
