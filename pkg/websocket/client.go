package websocket

import (
	"context"
	"errors"
	"github.com/adnanahmady/go-websocket-chat/pkg/request"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"io"
)

type Client struct {
	ctx  context.Context
	hub  *Hub
	conn *websocket.Conn
	send chan Message
	name string
}

func NewClient(
	ctx context.Context,
	hub *Hub,
	conn *websocket.Conn,
	name string,
) *Client {
	return &Client{
		ctx:  ctx,
		hub:  hub,
		send: make(chan Message, 100),
		conn: conn,
		name: name,
	}
}

func (c *Client) Register() {
	c.hub.register <- c
}

func (c *Client) Unregister() {
	lgr := request.GetLogger(c.ctx)

	msgType := websocket.CloseNormalClosure
	msg := websocket.FormatCloseMessage(msgType, "Server closed connection")
	if err := c.conn.WriteMessage(websocket.CloseMessage, msg); err != nil {
		lgr.Error("failed to send close message to client", "error", err)
		return
	}

	c.hub.unregister <- c

	if err := c.conn.Close(); err != nil {
		lgr.Error("failed to unregister client", "error", err)
	}
}

func (c *Client) Read() {
	lgr := request.GetLogger(c.ctx)

	defer func() {
		if r := recover(); r != nil {
			lgr.Error("recovered form panic: %v", r)
		}
		c.Unregister()
	}()

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if errors.Is(err, io.ErrUnexpectedEOF) {
				lgr.Debug("EOF: %+v", err)
				continue
			}

			var closeErr *websocket.CloseError
			if errors.As(err, &closeErr) {
				lgr.Error("close error occurred", "error", err)
				break
			}

			lgr.Error("failed to read message", "error", err)
			break
		}

		if msg.Payload == nil || msg.Payload == "" {
			continue
		}
		msg.ID = uuid.NewString()
		msg.Sender.Name = c.name

		c.hub.broadcast <- msg
	}
}

func (c *Client) Write() {
	lgr := request.GetLogger(c.ctx)

	defer func() {
		if r := recover(); r != nil {
			lgr.Error("recovered from panic when writing: %v", r)
		}
		c.Unregister()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					lgr.Error("failed to send close connection", "error", err)
				}
				lgr.Error("client channel named `send` is closed")
				return
			}
			if err := c.conn.WriteJSON(msg); err != nil {
				lgr.Error("error on writing message", "error", err)
			}
		case <-c.ctx.Done():
			lgr.Info("context done", "error", c.ctx.Err(), "section", "client")
			return
		}
	}
}
