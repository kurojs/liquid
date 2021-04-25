package main

import (
	"log"
	"net/http"

	"kurojs.github.com/liquid/auth_server"
	"kurojs.github.com/liquid/client"
	"kurojs.github.com/liquid/commons"
)

func main() {
	go func() {
		if err := client.Start(commons.GetInt("SERVER_PORT", 9090)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start client service failed: %s\n", err)
		}
	}()

	if err := auth_server.Start(commons.GetInt("CLIENT_PORT", 8090)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Start auth service failed: %s\n", err)
	}
}
