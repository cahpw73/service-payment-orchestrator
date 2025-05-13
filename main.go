package main

import (
	"log"
	"service-payment-orchestrator/app/factories"
	"service-payment-orchestrator/app/routes"
	"service-payment-orchestrator/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	log.Println("Loading configuration...")
	config.LoadConfig()

	appFactory, err := factories.NewFactory(config.RedisHost)

	if err != nil {
		log.Fatalf("Error occurred while creating the Factory: %v", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app, *appFactory)
	app.Listen(":" + config.Port)

	log.Println("Server running on port " + config.Port)
}
