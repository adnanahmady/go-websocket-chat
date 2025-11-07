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
	engine.POST("/ws", handler.WsEndpoint)
	engine.StaticFS("/static", http.Dir("static"))

	return &Routes{}
}
