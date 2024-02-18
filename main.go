package main

import (
	"go-on-store/app/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.SetupRouter()

	r.Run(":8080")
}
