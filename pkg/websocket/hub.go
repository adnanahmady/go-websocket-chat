package websocket

import (
	"context"
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"github.com/google/uuid"
	"net/http"

	"github.com/gorilla/websocket"
)

type Hub struct {
	logger     applog.Logger
	upgrader   *websocket.Upgrader
	broadcast  chan Message
	clients    map[*Client]*Client
	register   chan *Client
	unregister chan *Client
}

func NewHub(logger applog.Logger) *Hub {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &Hub{
		logger:     logger,
		upgrader:   upgrader,
		broadcast:  make(chan Message, 500),
		register:   make(chan *Client, 50),
		unregister: make(chan *Client, 50),
		clients:    make(map[*Client]*Client),
	}
}

func (h *Hub) RegisterClient(client *Client) {
	if _, exists := h.clients[client]; exists {
		return
	}
	h.clients[client] = client
	h.handleBroadcast(Message{
		ID:     uuid.NewString(),
		Type:   "join",
		Sender: Sender{Name: client.name},
	})
}

func (h *Hub) UnregisterClient(client *Client) {
	if _, exists := h.clients[client]; !exists {
		return
	}
	delete(h.clients, client)
	h.handleBroadcast(Message{
		ID:     uuid.NewString(),
		Type:   "leave",
		Sender: Sender{Name: client.name},
	})
}

func (h *Hub) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return h.upgrader.Upgrade(w, r, nil)
}

func (h *Hub) Run(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			h.logger.Error("recovered from panic on run: %v", r)
		}
	}()

	for {
		select {
		case client := <-h.register:
			h.RegisterClient(client)
		case client := <-h.unregister:
			h.UnregisterClient(client)
		case msg := <-h.broadcast:
			go h.handleBroadcast(msg)
		case <-ctx.Done():
			h.logger.Info("hub is stopped by the context", "reason", ctx.Err())
			return
		}
	}
}

func (h *Hub) handleBroadcast(msg Message) {
	if msg.Type != "join" && msg.Type != "leave" {
		msg.Type = "talk"
	}
	for _, c := range h.clients {
		if c.name == msg.Sender.Name {
			continue
		}
		c.send <- msg
	}
}
