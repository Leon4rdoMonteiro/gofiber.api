package main

import (
	"log"
	"os"

	"github.com/Leon4rdoMonteiro/gofiber.api/config"
	"github.com/Leon4rdoMonteiro/gofiber.api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to database
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}

	SERVER_PORT := os.Getenv("SERVER_PORT")

	// Create a Fiber application
	app := fiber.New()

	// Use fiber logger
	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":" + SERVER_PORT)
}
