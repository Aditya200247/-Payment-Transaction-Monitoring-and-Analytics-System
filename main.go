package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"payment-monitor/routes"
)

func main() {

	port := getEnv("PORT", "8080")

	gin.SetMode(getEnv("GIN_MODE", gin.DebugMode))

	log.Printf("Starting server | mode=%s port=%s", getEnv("GIN_MODE", gin.DebugMode), port)

	router := gin.Default()
	routes.SetupRoutes(router)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}