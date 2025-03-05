package chat

import (
	"fmt"
	"github.com/live-translate-edu/internal/configs"
	"github.com/live-translate-edu/internal/dto"
)

type Hub struct {
	register   chan *Client
	unregister chan *Client
	broadcast  chan *dto.MessageDto
	clients    map[string]map[*Client]bool
	cfg        *configs.Config
	handlers   *MessageHandlers
}

func NewHub(handlers *MessageHandlers) *Hub {
	return &Hub{
		broadcast:  make(chan *dto.MessageDto),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
		handlers:   handlers,
	}
}

func (h *Hub) AddRegisterClient(client *Client) {
	h.register <- client
}

func (h *Hub) RegisterClient(client *Client) {
	if connect, ok := h.clients[client.room]; !ok {
		connect = make(map[*Client]bool)
		h.clients[client.room] = connect
	}
	h.clients[client.room][client] = true
	fmt.Println(fmt.Sprintf("Количество клиентов в комнате %s: %d", client.room, len(h.clients[client.room])))
}

func (h *Hub) UnregisterClient(client *Client) {
	if _, ok := h.clients[client.room]; ok {
		delete(h.clients[client.room], client)
		close(client.send)
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.RegisterClient(client)
		case client := <-h.unregister:
			h.UnregisterClient(client)
		case message := <-h.broadcast:
			h.Broadcast(message)
		}
	}
}

func (h *Hub) Broadcast(message *dto.MessageDto) {
	clients := h.clients[message.Room]
	h.handlers.HandleMessage(message)
	for client := range clients {
		select {
		case client.send <- message:
		default:
			h.unregister <- client
		}
	}
}
