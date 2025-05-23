package main

import (
	"gateway/routes"
	"log"
	"os"
	dotenv "github.com/joho/godotenv"
)

func main() {
	// Set log output to file
	file, err := os.Create("/logs/gateway.log")
	if err != nil {
		log.Println("Failed setting log output to file")
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("Starting Gateway")

	if err := dotenv.Load(); err != nil {
		log.Println("Couldn't load .env file, continuing with default values")
	}
	routes.Start()
}