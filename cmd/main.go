package main

import (
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/config"
	"github.com/abdullahnettoor/connect-social-media/internal/di"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(c)
	}

	server, err := di.InitializeAPI(c)
	if err != nil {
		log.Fatalln(c)
	}

	server.Start()
}
