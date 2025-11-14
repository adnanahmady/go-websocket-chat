package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ps, err := WireUpApp()
	if err != nil {
		log.Fatalf("failed to wire up the application dependencies")
	}

	ps.Logger.Info("App Name: %s", ps.Config.App.Name)
	ps.Logger.Info(
		"Application port is correctly set",
		"port", ps.Config.App.Port,
	)
	ps.Logger.Info("Host: %s", ps.Config.App.Host)
	ps.Logger.Info("Env: %s", ps.Config.App.Env)
	go runServer(ps)
	go ps.Hub.Run(ctx)

	<-ctx.Done()
	stop()

	if err := ps.Server.Shutdown(); err != nil {
		os.Exit(1)
	}
	ps.Logger.Info("Application stopped")
}

func runServer(ps *App) {
	if err := ps.Server.Start(); err != nil {
		os.Exit(1)
	}
}
