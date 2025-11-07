package presentation

import (
	"fmt"
	"net/http"

	"github.com/adnanahmady/go-websocket-chat/pkg/request"
	"github.com/adnanahmady/go-websocket-chat/pkg/websocket"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HomePage(c *gin.Context) {
	_, _ = fmt.Fprintf(c.Writer, "Welcome to the HomePage!")
}

func (h *Handler) WsEndpoint(c *gin.Context) {
	lgr := request.GetLogger(c.Request.Context())
	lgr.Info("upgrading to websocket")
	websocket.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		lgr.Error("failed to upgrade to websocket", "error", err)
		return
	}

	lgr.Info("client successfully connected...")
	websocket.Reader(c.Request.Context(), ws)
}
