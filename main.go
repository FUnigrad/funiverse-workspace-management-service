package main

import (
	"log"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/handler"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Cannot load config: ", err)
	}

	server := handler.NewServer(config)

	err = server.Start()
	if err != nil {
		log.Fatalln("Cannot start server:", err)
	}
}
