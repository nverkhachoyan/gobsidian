package main

import (
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/parsers"
	"gobsidian/internal/renderers"
	"gobsidian/internal/transformers"

	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

type App struct {
	logger *log.Logger
	cfg    config.Config
	gen    *transformers.SiteTransformer
	hub    *Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewApp() (*App, error) {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
		Prefix:          "Cooking 🍪",
	})

	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		return nil, err
	}

	markdownRenderer := renderers.NewMarkdownRenderer()
	parser := parsers.NewObsidianParser(cfg)

	gen, err := transformers.NewSiteTransformer(
		cfg,
		logger,
		parser,
		markdownRenderer,
	)
	if err != nil {
		return nil, err
	}

	hub := newHub()
	go hub.run()

	return &App{
		logger: logger,
		cfg:    cfg,
		gen:    gen,
		hub:    hub,
	}, nil
}

func (a *App) Run(clear, watch, serve bool, port string) {
	if clear {
		if err := os.RemoveAll(a.cfg.OutputDirectory); err != nil {
			a.logger.Error("Failed to clear output directory", "error", err)
			os.Exit(1)
		}
	}

	if err := a.gen.GenerateSite(); err != nil {
		a.logger.Error("Failed to generate site", "error", err)
		os.Exit(1)
	}

	if watch {
		if serve {
			a.logger.Print("Watching for changes and serving...", "port", port)
			go a.serve(port)
		} else {
			a.logger.Print("Watching for changes...")
		}
		a.watch()
	} else if serve {
		a.logger.Print("Building complete. Serving website...", "port", port)
		a.serve(port)
	} else {
		a.logger.Print("Build complete. Use --watch for live reloading or --serve to host.")
	}
}

func (a *App) serve(port string) {
	fs := http.FileServer(http.Dir(a.cfg.OutputDirectory))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			a.logger.Error("Failed to upgrade websocket", "error", err)
			return
		}
		a.hub.register <- conn
	})

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(a.cfg.OutputDirectory, r.URL.Path)

		stat, err := os.Stat(requestedPath)
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, filepath.Join(a.cfg.OutputDirectory, "404.html"))
			return
		}

		if stat.IsDir() {
			indexPath := filepath.Join(requestedPath, "index.html")
			if _, err := os.Stat(indexPath); os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, filepath.Join(a.cfg.OutputDirectory, "404.html"))
				return
			}
		}
		fs.ServeHTTP(w, r)
	}))

	addr := ":" + port
	a.logger.Print("Serving static website", "url", "http://localhost"+addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		a.logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

// watch monitors the filesystem for changes and triggers a site rebuild.
func (a *App) watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		a.logger.Error("Failed to create file watcher", "error", err)
		os.Exit(1)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) != 0 {
					a.logger.Infof("Change detected in %s, regenerating site", event.Name)
					if err := a.gen.GenerateSite(); err != nil {
						a.logger.Error("Error regenerating site", "error", err)
					} else {
						// On successful regeneration, notify clients to reload.
						a.hub.broadcast <- []byte("reload")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				a.logger.Error("Watcher error", "error", err)
			}
		}
	}()

	// Add non-recursive paths first
	otherPaths := []string{"templates", "config.yaml"}
	for _, path := range otherPaths {
		fi, err := os.Stat(path)
		if err != nil {
			a.logger.Warn("Could not stat path, skipping watch", "path", path, "error", err)
			continue
		}
		if fi.IsDir() {
			a.logger.Print("Watching for changes in directory", "path", path)
			if err := watcher.Add(path); err != nil {
				a.logger.Error("Failed to watch directory", "directory", path, "error", err)
			}
		} else {
			a.logger.Print("Watching for changes in file", "path", path)
			if err := watcher.Add(path); err != nil {
				a.logger.Error("Failed to watch file", "file", path, "error", err)
			}
		}
	}

	// Recursively watch the input directory
	err = filepath.WalkDir(a.cfg.InputDirectory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if err := watcher.Add(path); err != nil {
				a.logger.Error("Failed to watch directory", "directory", path, "error", err)
			}
		}
		return nil
	})
	if err != nil {
		a.logger.Error("Error walking input directory", "error", err)
	}

	// Wait for a shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.logger.Info("Shutdown signal received, exiting gracefully.")
}
