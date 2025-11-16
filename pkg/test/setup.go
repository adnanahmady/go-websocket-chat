package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adnanahmady/go-websocket-chat/internal"
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"strings"
	"testing"
	"time"
)

type WebsocketServer struct {
	server     *httptest.Server
	App        *internal.App
	CancelFunc context.CancelFunc
	wsUrl      string
}

func Setup() (*WebsocketServer, error) {
	app, err := internal.WireUpApp()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go app.Hub.Run(ctx)
	srv := httptest.NewServer(app.Server.GetEngine())
	wsUrl := fmt.Sprintf(
		"ws%s/ws",
		strings.TrimPrefix(srv.URL, "http"),
	)
	return &WebsocketServer{
		server:     srv,
		App:        app,
		CancelFunc: cancel,
		wsUrl:      wsUrl,
	}, nil
}

func (w *WebsocketServer) Close() {
	w.server.Close()
}

func (w *WebsocketServer) Connect(user string) (*Connection, error) {
	headers := make(http.Header)
	url := fmt.Sprintf("%s?sender=%s", w.wsUrl, user)
	conn, res, err := websocket.DefaultDialer.Dial(url, headers)
	if err != nil {
		return nil, fmt.Errorf("fialed to dial websocket server (%w)", err)
	}
	go func() {
		w.App.Logger.Info("closing websocket connection in 10 seconds")
		time.Sleep(10 * time.Second)
		if err := conn.Close(); err != nil {
			w.App.Logger.Error("failed to close websocket connection", "error", err)
		}
	}()
	if err := conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline (%w)", err)
	}
	return &Connection{conn: conn, ConnectResult: res}, nil
}

type Connection struct {
	conn          *websocket.Conn
	ConnectResult *http.Response
}

func (c *Connection) SetReadDeadline(d time.Duration) error {
	if err := c.conn.SetReadDeadline(time.Now().Add(d * time.Millisecond)); err != nil {
		return fmt.Errorf("failed to set read deadline (%w)", err)
	}
	return nil
}

func (c *Connection) Send(msg any) error {
	jm, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal websocket message (%w)", err)
	}
	if err := c.conn.WriteMessage(websocket.TextMessage, jm); err != nil {
		return fmt.Errorf("failed to send message (%w)", err)
	}
	return nil
}

type Message map[string]any

func (c *Connection) Read() (Message, error) {
	var msg Message
	err := c.conn.ReadJSON(&msg)
	if err != nil {
		return nil, fmt.Errorf("failed to read message (%w)", err)
	}
	return msg, nil
}

// ReadType read the message type or return deadline error
func (c *Connection) ReadType(msgType string) (Message, error) {
	for {
		msg, err := c.Read()
		if err != nil || msg["type"] == msgType {
			return msg, err
		}
	}
	// this will never get executed because the
	// read operation eventually faces deadline error
}

// ShouldReadType it reads the message with expected type or test fails
func (c *Connection) ShouldReadType(t testing.TB, msgType string) Message {
	msg, err := c.ReadType(msgType)
	if err != nil {
		t.Fatal(err, string(debug.Stack()))
	}
	return msg
}

func (c *Connection) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection (%w)", err)
	}
	return nil
}
