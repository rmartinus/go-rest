package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rmartinus/go-rest/pkg/api"
	"github.com/rmartinus/go-rest/pkg/logger"
)

func main() {
	log := logger.CreateLogger("go-rest")

	router, err := api.NewRouter(api.Handlers(), "./v1.yaml")
	if err != nil {
		log.Panicf("Router creation error: %s", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Infof("Listening on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
