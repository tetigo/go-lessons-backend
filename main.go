package main

import (
	"log"

	"github.com/tetigo/go-lessons-backend/internal/server"
)

func main() {
	myServer,err := server.New(server.Config{Port: 9090})
	if err !=nil{
		log.Fatalf("impossible to create the server: %s", err)
	}

	myServer.Run() // runs on 0.0.0.0:9090
}

// linux: sudo lfof -i -P | grep LISTEN
// windows: netstat -ab
