package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gocorreio.rastreio/config"
	handler "github.com/jeffotoni/gocorreio.rastreio/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v2/", handler.Rastreio2)
	mux.HandleFunc("/api/v2", handler.Rastreio2)
	mux.HandleFunc("/api/v1/", handler.Rastreio)
	mux.HandleFunc("/api/v1", handler.NotFound)
	mux.HandleFunc("/", handler.NotFound)

	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	log.Println("Port:", config.Port)
	log.Fatal(server.ListenAndServe())
}
