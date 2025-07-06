package server

import (
	"context"
	"gobsidian/internal/websockets"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

type WebServer struct {
	outputDirectory string
	logger          *log.Logger
	port            string
	server          *http.Server
	serveMux        *http.ServeMux
	hub             *websockets.Hub
}

func NewWebServer(port string, outputDirectory string, logger *log.Logger) *WebServer {
	mux := http.NewServeMux()

	hub := websockets.NewHub(logger)
	go hub.Run()

	return &WebServer{
		outputDirectory: outputDirectory,
		logger:          logger,
		port:            port,
		serveMux:        mux,
		hub:             hub,
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (ws *WebServer) Serve() {
	ws.serveMux.HandleFunc("/ws", ws.handleWebSocket)
	ws.serveMux.Handle("/", ws.fileServerHandler())

	ws.logger.Info("Serving static website", "url", "http://localhost:"+ws.port)

	go func() {
		if err := ws.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ws.logger.Error("HTTP server failed to start or shut down unexpectedly", "error", err)
			os.Exit(1)
		}
		ws.logger.Debug("HTTP server stopped.")
	}()

	time.Sleep(1000 * time.Millisecond)
	ws.hub.Broadcast <- []byte("reload")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ws.logger.Info("Shutdown signal received, exiting gracefully.")
	ws.server.Shutdown(context.Background())
}

func (ws *WebServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
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
		ws.logger.Error("Failed to upgrade websocket connection", "error", err)
		return
	}
	ws.hub.Register <- conn
}

func (ws *WebServer) fileServerHandler() http.HandlerFunc {
	fs := http.FileServer(http.Dir(ws.outputDirectory))
	return func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(ws.outputDirectory, r.URL.Path)

		stat, err := os.Stat(requestedPath)
		if os.IsNotExist(err) {
			ws.logger.Warn("File not found, serving 404", "path", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, filepath.Join(ws.outputDirectory, "404.html"))
			return
		}
		if err != nil {
			ws.logger.Error("Error stating file", "path", r.URL.Path, "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if stat.IsDir() {
			indexPath := filepath.Join(requestedPath, "index.html")
			if _, err := os.Stat(indexPath); os.IsNotExist(err) {
				ws.logger.Warn("Directory without index.html, serving 404", "path", r.URL.Path)
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, filepath.Join(ws.outputDirectory, "404.html"))
				return
			}
		}
		fs.ServeHTTP(w, r)
	}
}
