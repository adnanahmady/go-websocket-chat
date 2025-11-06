//go:build wireinject
// +build wireinject

package main

import (
	"github.com/adnanahmady/go-websocket-chat/config"
	"github.com/adnanahmady/go-websocket-chat/pkg/applog"
	"github.com/google/wire"
)

type App struct {
	Config *config.Config
	Logger *applog.AppLogger
}

var AppSet = wire.NewSet(
	applog.NewAppLogger,
	wire.Bind(new(applog.Logger), new(*applog.AppLogger)),

	config.GetConfig,

	wire.Struct(new(App), "*"),
)

func WireUpApp() (*App, error) {
	wire.Build(AppSet)
	return nil, nil
}
