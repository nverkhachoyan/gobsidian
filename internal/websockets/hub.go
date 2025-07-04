package websockets

import (
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan []byte
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
	logger     *log.Logger
}

func NewHub(logger *log.Logger) *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		Clients:    make(map[*websocket.Conn]bool),
		logger:     logger,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.logger.Debug("Client registered", "remote_addr", client.RemoteAddr())
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				_ = client.Close()
				h.logger.Debug("Client unregistered", "remote_addr", client.RemoteAddr())
			}
		case message := <-h.Broadcast:
			h.logger.Debug("Broadcasting message", "message", string(message), "clients", len(h.Clients))
			for client := range h.Clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					h.logger.Error("Failed to write message to client", "error", err, "remote_addr", client.RemoteAddr())
					_ = client.Close()
					delete(h.Clients, client)
				}
			}
		}
	}
}
