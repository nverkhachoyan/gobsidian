package main

import (
	"flag"

	"os"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/generators"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/renderers"
	"gobsidian/internal/server"

	"github.com/charmbracelet/log"
)

func main() {
	clearFlag := flag.Bool("clear", false, "Clear the public folder before generating")
	serveFlag := flag.Bool("serve", false, "Serve the generated website after building")
	portFlag := flag.String("port", "8000", "Port to serve the website on")
	vaultHealthFlag := flag.Bool("vhealth", false, "Check the health of the vault")
	flag.Parse()

	cfg, err := config.ReadConfig(".config.toml")
	if err != nil {
		log.Fatal(err)
	}

	loggerLevel := log.InfoLevel
	if cfg.Env == "development" {
		loggerLevel = log.DebugLevel
	}

	logger := log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
		Level:           loggerLevel,
	})

	markdownRenderer := renderers.NewMarkdownRenderer()
	parser := parsers.NewObsidianParser(&parsers.ObsidianParser{
		InputDirectory: cfg.InputDirectory,
		Regexes: &models.ParserRegexes{
			FrontmatterRegex:   cfg.RegexpConfig.FrontmatterRegex,
			WikilinkRegex:      cfg.RegexpConfig.WikilinkRegex,
			ObsidianImageRegex: cfg.RegexpConfig.ObsidianImageRegex,
			HashtagRegex:       cfg.RegexpConfig.HashtagRegex,
		},
	})

	gen, err := generators.NewStaticSiteGenerator(
		cfg.InputDirectory,
		cfg.OutputDirectory,
		cfg.SiteTitle,
		cfg.SiteSubtitle,
		cfg.BaseURL,
		cfg.NotesPerPage,
		&models.ParserRegexes{
			FrontmatterRegex:   cfg.RegexpConfig.FrontmatterRegex,
			WikilinkRegex:      cfg.RegexpConfig.WikilinkRegex,
			ObsidianImageRegex: cfg.RegexpConfig.ObsidianImageRegex,
			HashtagRegex:       cfg.RegexpConfig.HashtagRegex,
		},
		logger,
		cfg.Templates,
		parser,
		markdownRenderer,
		*vaultHealthFlag,
	)

	if err != nil {
		log.Error("Failed to initialize application", "error", err)
		os.Exit(1)
	}

	if *clearFlag {
		if err := os.RemoveAll(cfg.OutputDirectory); err != nil {
			logger.Error("Failed to clear output directory", "error", err)
			os.Exit(1)
		}
	}

	if err := gen.Generate(); err != nil {
		logger.Error("Failed to generate site", "error", err)
		os.Exit(1)
	}

	if *serveFlag {
		webServer := server.NewWebServer(*portFlag, cfg.OutputDirectory, logger)
		logger.Info("Serving website...", "port", *portFlag)
		webServer.Serve()
	}
}
