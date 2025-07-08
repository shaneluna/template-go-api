package main

import (
	"log"
	"myproject/internal/app"
)

func main() {
	router, err := app.NewGinApp()
	if err != nil {
		log.Fatalf("Failed to create gin app: %v", err)
	}

	log.Println("Starting gin server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Gin server failed: %v", err)
	}
} 