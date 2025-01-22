package main

import (
	"context"
	"log"
	"rendipilot/simple-golang/database"
	"rendipilot/simple-golang/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection ONCE
	_, err = database.ConnectDatabase() 
	if err != nil { 
		log.Fatal("Failed to connect to PostgreSQL:", err) 
	} 

	defer database.GetDB().Close(context.Background()) 

	// Define endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello Rendy")
	})

	app.Post("/adduser", handlers.AddUser)

	// Start the server
	app.Listen(":3000")
}
