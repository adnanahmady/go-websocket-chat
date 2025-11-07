//go:build wireinject
// +build wireinject

package main

import (
	"github.com/adnanahmady/go-websocket-chat/config"
	"github.com/adnanahmady/go-websocket-chat/internal/presentation"
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"github.com/adnanahmady/go-websocket-chat/pkg/request"
	"github.com/google/wire"
)

type App struct {
	Config *config.Config
	Logger *applog.AppLogger
	Server *request.Server
	Routes *presentation.Routes
}

var AppSet = wire.NewSet(
	applog.NewAppLogger,
	wire.Bind(new(applog.Logger), new(*applog.AppLogger)),
	config.GetConfig,
	request.NewServer,
	wire.Bind(new(request.Router), new(*request.Server)),

	presentation.NewHandler,
	presentation.NewRoutes,

	wire.Struct(new(App), "*"),
)

func WireUpApp() (*App, error) {
	wire.Build(AppSet)
	return nil, nil
}
