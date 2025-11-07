package websocket

import (
	"context"

	"github.com/adnanahmady/go-websocket-chat/pkg/request"
	"github.com/gorilla/websocket"
)

var Upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Reader(ctx context.Context, conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			request.GetLogger(ctx).Error("failed to read message", "error", err)
			return
		}

		request.GetLogger(ctx).Info(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			request.GetLogger(ctx).Error("failed to write message", "error", err)
			return
		}
	}
}
