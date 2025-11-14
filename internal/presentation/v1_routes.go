package presentation

import (
	"github.com/adnanahmady/go-websocket-chat/pkg/request"
	"net/http"
)

type Routes struct{}

func NewRoutes(
	router request.Router,
	handler *Handler,
) *Routes {
	engine := router.GetEngine()
	engine.GET("/", handler.HomePage)
	engine.GET("/ws", handler.WsEndpoint)
	engine.StaticFS("/chat", http.Dir("static"))

	return &Routes{}
}
