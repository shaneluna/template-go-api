package main

import (
	"log"
	"myproject/internal/app"
	"net/http"
)

func main() {
	router, err := app.NewApp()
	if err != nil {
		log.Fatalf("Failed to create chi app: %v", err)
	}

	log.Println("Starting chi server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Chi server failed: %v", err)
	}
}
