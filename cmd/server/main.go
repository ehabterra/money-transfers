package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	s := prepareRoutes()

	listenURL := fmt.Sprintf("http://localhost:%v", port)

	fmt.Printf("Listening on %v\n", listenURL)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), s))
}
