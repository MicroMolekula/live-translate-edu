package chat

import (
	"github.com/gorilla/websocket"
	"github.com/live-translate-edu/internal/dto"
	"log"
)

type Client struct {
	room string
	conn *websocket.Conn
	send chan *dto.MessageDto
	hub  *Hub
	user *dto.UserDTO
}

func NewClient(room string, conn *websocket.Conn, hub *Hub, user *dto.UserDTO) *Client {
	return &Client{
		room: room,
		conn: conn,
		send: make(chan *dto.MessageDto, 256),
		hub:  hub,
		user: user,
	}
}

func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
		if err := c.conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		var responseMessage *dto.ResponseMessageDto
		if err := c.conn.ReadJSON(&responseMessage); err != nil {
			log.Println(err)
			break
		}
		c.hub.broadcast <- &dto.MessageDto{User: c.user, Content: responseMessage.Content, Room: c.room}
	}
}

func (c *Client) Write() {
	defer func() {
		if err := c.conn.Close(); err != nil {
			log.Println(err)
		}
	}()
	for m := range c.send {
		if err := c.conn.WriteJSON(m); err != nil {
			log.Println(err)
		}
	}
}

func (c *Client) Close() {
	close(c.send)
}

func (c *Client) getUser() *dto.UserDTO {
	return c.user
}
