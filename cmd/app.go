package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/generators"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/renderers"
	"gobsidian/internal/websockets"

	"github.com/charmbracelet/log"
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

	parser := parsers.NewObsidianParser(&parsers.ObsidianParser{
		InputDirectory: cfg.InputDirectory,
		Regexes: &models.ParserRegexes{
			FrontmatterRegex:   cfg.RegexpConfig.FrontmatterRegex,
			WikilinkRegex:      cfg.RegexpConfig.WikilinkRegex,
			ObsidianImageRegex: cfg.RegexpConfig.ObsidianImageRegex,
			HashtagRegex:       cfg.RegexpConfig.HashtagRegex,
		},
	})

	gen, err := generators.NewStaticSiteGenerator(&generators.StaticSiteGenerator{
		InputDirectory:  cfg.InputDirectory,
		OutputDirectory: cfg.OutputDirectory,
		SiteTitle:       cfg.SiteTitle,
		SiteSubtitle:    cfg.SiteSubtitle,
		BaseURL:         cfg.BaseURL,
		Templates:       cfg.Templates,
		Regexes: &models.ParserRegexes{
			FrontmatterRegex:   cfg.RegexpConfig.FrontmatterRegex,
			WikilinkRegex:      cfg.RegexpConfig.WikilinkRegex,
			ObsidianImageRegex: cfg.RegexpConfig.ObsidianImageRegex,
			HashtagRegex:       cfg.RegexpConfig.HashtagRegex,
		},
		Logger:           logger,
		Parser:           parser,
		MarkdownRenderer: markdownRenderer,
	})
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
