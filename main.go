package main

import (
	"go-rest-api/config"
	"os"
)

const (
	PORT = "9000"
)

func main() {
	//Get port from command line interface
	port := os.Getenv("Port")
	if port == "" {
		port = PORT
	}
	app := &config.App{}
	app.Intialize()
	app.Run(port)
}
