package main

import (
	"io"
	"log"
	"net/http"

	goclient "github.com/FUnigrad/funiverse-workspace-service/internal/goclient"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/pods", goclient.GetPods)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
