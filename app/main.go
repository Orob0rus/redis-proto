package main

import (
	"fmt"
	"os"
	"os/signal"
	"redis/app/server"
	"syscall"
)

// TODO: Add golang profiler
func main() {
	mayday := make(chan os.Signal, 1)
	signal.Notify(mayday, syscall.SIGINT, syscall.SIGTERM)

	server := &server.Server{
		Handler: server.Handler{},
	}
	server.Init()

	<-mayday
	fmt.Println("Shutting down the server")
}
