package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/generators"
	"gobsidian/internal/parsers"
	"gobsidian/internal/renderers"
	"gobsidian/internal/websockets"

	"github.com/charmbracelet/log"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

type App struct {
	logger *log.Logger
	cfg    config.Config
	gen    *generators.StaticSiteGenerator
	hub    *websockets.Hub
	server *http.Server
}

func NewApp() (*App, error) {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
		Prefix:          "Cooking 🍪",
		Level:           log.DebugLevel,
	})

	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		return nil, err
	}

	markdownRenderer := renderers.NewMarkdownRenderer()
	parser := parsers.NewObsidianParser(cfg)

	gen, err := generators.NewStaticSiteGenerator(
		cfg,
		logger,
		parser,
		markdownRenderer,
	)
	if err != nil {
		return nil, err
	}

	hub := websockets.NewHub(logger)
	go hub.Run()

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

	if serve {
		a.logger.Print("Serving website...", "port", port)
		a.serve(port)
	}

	// if watch {
	// 	if serve {
	// 		a.logger.Print("Watching for changes and serving...", "port", port)
	// 		go a.serve(port)
	// 	} else {
	// 		a.logger.Print("Watching for changes...")
	// 	}
	// 	a.watch()
	// } else if serve {
	// 	a.logger.Print("Building complete. Serving website...", "port", port)
	// 	a.serve(port)
	// } else {
	// 	a.logger.Print("Build complete. Use --watch for live reloading or --serve to host.")
	// }
}

func (a *App) serve(port string) {
	mux := http.NewServeMux()

	a.server = &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	mux.HandleFunc("/ws", a.handleWebSocket)
	mux.Handle("/", a.fileServerHandler())

	a.logger.Print("Serving static website", "url", "http://localhost:"+port)

	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Error("HTTP server failed to start or shut down unexpectedly", "error", err)
			os.Exit(1)
		}
		a.logger.Debug("HTTP server stopped.")
	}()

	time.Sleep(1000 * time.Millisecond)
	a.hub.Broadcast <- []byte("reload")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.logger.Info("Shutdown signal received, exiting gracefully.")
	a.server.Shutdown(context.Background())
}

func (a *App) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// TODO: Add proper origin checking
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.logger.Error("Failed to upgrade websocket connection", "error", err)
		return
	}
	a.hub.Register <- conn
}

func (a *App) fileServerHandler() http.HandlerFunc {
	fs := http.FileServer(http.Dir(a.cfg.OutputDirectory))
	return func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(a.cfg.OutputDirectory, r.URL.Path)

		stat, err := os.Stat(requestedPath)
		if os.IsNotExist(err) {
			a.logger.Warn("File not found, serving 404", "path", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, filepath.Join(a.cfg.OutputDirectory, "404.html"))
			return
		}
		if err != nil {
			a.logger.Error("Error stating file", "path", r.URL.Path, "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if stat.IsDir() {
			indexPath := filepath.Join(requestedPath, "index.html")
			if _, err := os.Stat(indexPath); os.IsNotExist(err) {
				a.logger.Warn("Directory without index.html, serving 404", "path", r.URL.Path)
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, filepath.Join(a.cfg.OutputDirectory, "404.html"))
				return
			}
		}
		fs.ServeHTTP(w, r)
	}
}

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
				a.logger.Debug("Watcher event", "event", event)

				if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename|fsnotify.Remove|fsnotify.Chmod) != 0 {
					a.logger.Infof("Change detected in %s, regenerating site", event.Name)
					if err := os.RemoveAll(a.cfg.OutputDirectory); err != nil {
						a.logger.Error("Failed to clear output directory", "error", err)
						os.Exit(1)
					}
					if err := a.gen.GenerateSite(); err != nil {
						a.logger.Error("Error regenerating site", "error", err)
					} else {
						a.hub.Broadcast <- []byte("reload")
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

	watchPaths := []string{"config.yaml"}
	if err := a.addWatchFiles(watcher, watchPaths...); err != nil {
		a.logger.Error("Failed to watch files", "error", err)
	}

	watchDirs := []string{a.cfg.InputDirectory, "templates"}
	if err := a.addRecursiveWatchDirs(watcher, watchDirs...); err != nil {
		a.logger.Error("Failed to watch directories", "error", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.logger.Info("Shutdown signal received, exiting gracefully.")
}

func (a *App) addRecursiveWatchDirs(watcher *fsnotify.Watcher, dirs ...string) error {
	for _, dir := range dirs {
		a.logger.Info("Adding recursive watch for directory", "directory", dir)
		err := filepath.WalkDir(dir, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr // Propagate walk errors
			}
			// Add both directories and files to the watcher
			if err := watcher.Add(path); err != nil {
				a.logger.Error("Failed to add path to watcher", "path", path, "error", err)
				// Don't exit here, try to continue with other paths
			} else {
				a.logger.Debug("Successfully added path to watcher", "path", path)
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("error walking directory %s: %w", dir, err) // Return error from walk
		}
	}
	return nil
}

func (a *App) addWatchFiles(watcher *fsnotify.Watcher, paths ...string) error {
	for _, path := range paths {
		a.logger.Info("Adding watch for path", "path", path)
		fi, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				a.logger.Warn("Path does not exist, skipping watch", "path", path)
			} else {
				a.logger.Warn("Could not stat path, skipping watch", "path", path, "error", err)
			}
			continue
		}

		if fi.IsDir() {
			if err := watcher.Add(path); err != nil {
				a.logger.Error("Failed to add directory to watcher", "directory", path, "error", err)
			} else {
				a.logger.Debug("Successfully added directory to watcher", "directory", path)
			}
		} else {
			if err := watcher.Add(path); err != nil {
				a.logger.Error("Failed to add file to watcher", "file", path, "error", err)
			} else {
				a.logger.Debug("Successfully added file to watcher", "file", path)
			}
		}
	}
	return nil
}
