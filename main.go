package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

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

	<-termChan
	signal.Stop(termChan)
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
