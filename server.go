package main

import (
	"log"
	"os"

	"github.com/Leon4rdoMonteiro/gofiber.api/config"
	"github.com/Leon4rdoMonteiro/gofiber.api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}

	SERVER_PORT := os.Getenv("SERVER_PORT")

	// Create a Fiber application
	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":" + SERVER_PORT)
}
