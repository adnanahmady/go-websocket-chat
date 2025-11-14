package presentation

import (
	"fmt"
	"github.com/adnanahmady/go-websocket-chat/pkg/request"
	"github.com/adnanahmady/go-websocket-chat/pkg/websocket"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	hub *websocket.Hub
}

func NewHandler(hub *websocket.Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

func (h *Handler) HomePage(c *gin.Context) {
	_, _ = fmt.Fprintf(c.Writer, "Welcome to the Websocket chat app!")
}

func (h *Handler) WsEndpoint(c *gin.Context) {
	ctx := c.Request.Context()
	lgr := request.GetLogger(ctx)
	lgr.Info("upgrading to websocket")

	conn, err := h.hub.Upgrade(c.Writer, c.Request)
	if err != nil {
		lgr.Error("failed to upgrade to websocket", "error", err)
		return
	}

	name := c.Query("sender")
	if name == "" {
		name = "anonymouse"
	}
	wsCtx := request.ToWsCtx(ctx)
	client := websocket.NewClient(wsCtx, h.hub, conn, name)
	client.Register()

	go client.Read()
	go client.Write()

	lgr.Info("client successfully connected...")
}
