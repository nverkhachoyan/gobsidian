package main

import (
	"flag"
	"os"
	"time"

	"github.com/charmbracelet/log"

	"gobsidian/internal/config"
	"gobsidian/internal/executors"
	"gobsidian/internal/parsers"
	"gobsidian/internal/transformers"
	"gobsidian/internal/utils"

	"github.com/fsnotify/fsnotify"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat: time.DateTime,
	})

	watchFlag := flag.Bool("watch", false, "Watch for file changes and rebuild automatically")
	clearFlag := flag.Bool("clear", false, "Clear the public folder before generating")
	flag.Parse()

	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		logger.Error("Failed to read config", "error", err)
		os.Exit(1)
	}

	postExecutor, err := executors.NewPostExecutor(cfg)
	if err != nil {
		logger.Error("Failed to initialize post executor", "error", err)
		os.Exit(1)
	}
	indexExecutor, err := executors.NewIndexExecutor(cfg)
	if err != nil {
		logger.Error("Failed to initialize index executor", "error", err)
		os.Exit(1)
	}
	tagExecutor, err := executors.NewTagExecutor(cfg)
	if err != nil {
		logger.Error("Failed to initialize tag executor", "error", err)
		os.Exit(1)
	}
	previewExecutor, err := executors.NewPreviewExecutor(cfg)
	if err != nil {
		logger.Error("Failed to initialize preview executor", "error", err)
		os.Exit(1)
	}
	folderExecutor, err := executors.NewFolderExecutor(cfg)
	if err != nil {
		logger.Error("Failed to initialize folder executor", "error", err)
		os.Exit(1)
	}

	markdownRenderer := utils.NewMarkdownRenderer()
	parser := parsers.NewMarkdownParser(cfg)

	gen, err := transformers.NewSiteTransformer(
		cfg,
		logger,
		parser,
		markdownRenderer,
		postExecutor,
		indexExecutor,
		tagExecutor,
		previewExecutor,
		folderExecutor,
	)
	if err != nil {
		logger.Error("Failed to initialize generator", "error", err)
		os.Exit(1)
	}

	if *clearFlag {
		err := os.RemoveAll(cfg.OutputDirectory)
		if err != nil {
			logger.Error("Failed to clear output directory", "error", err)
			os.Exit(1)
		}
	}

	if err := gen.GenerateSite(); err != nil {
		logger.Error("Failed to generate site", "error", err)
		os.Exit(1)
	}

	if *watchFlag {
		watch(gen, cfg, logger)
	} else {
		logger.Info("Build complete. Use --watch for live reloading.")
	}
}

func watch(gen *transformers.SiteTransformer, cfg config.Config, logger *log.Logger) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error("Failed to create file watcher", "error", err)
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					logger.Infof("Change detected in %s, regenerating site", event.Name)
					if err := gen.GenerateSite(); err != nil {
						logger.Error("Error regenerating site", "error", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error("Watcher error", "error", err)
			}
		}
	}()

	watchDirs := []string{cfg.InputDirectory, "templates"}
	for _, dir := range watchDirs {
		logger.Infof("Watching for changes in '%s' directory", dir)
		if err := watcher.Add(dir); err != nil {
			logger.Error("Failed to watch directory", "error", err)
		}
	}

	<-done
}
