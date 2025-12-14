package main

import (
	"log"

	"github.com/eldanielhumberto/ct/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	// Execute command
	cmd.Execute()
}
