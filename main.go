package main

import (
	"fmt"
	"github.com/adnanahmady/go-websocket-chat/config"
)

func main() {
	cfg := config.GetConfig()

	fmt.Printf("App Name: %s\n", cfg.App.Name)
	fmt.Printf("Port: %d\n", cfg.App.Port)
	fmt.Printf("Host: %s\n", cfg.App.Host)
	fmt.Printf("Env: %s\n", cfg.App.Env)
}
