package main

import (
	"gateway/routes"
	"log"
)

func main() {
	log.Println("Starting Gateway")
	routes.Start()
}