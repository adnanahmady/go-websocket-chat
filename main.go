package main

import (
	"log"
)

func main() {
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
}
