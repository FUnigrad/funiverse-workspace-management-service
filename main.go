package main

import (
	"log"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/goclient"
	"github.com/FUnigrad/funiverse-workspace-service/handler"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	client, err := goclient.NewClient(config)
	if err != nil {
		log.Fatal("Cannot init K8s Client")
	}
	// // pod, _ := client.GetPodsName()

	// fmt.Printf("%s", &pod[1])

	server := handler.NewServer(client)

	err = server.Start(config)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
