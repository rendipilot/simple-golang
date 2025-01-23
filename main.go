package main

import (
	"context"
	"log"
	"rendipilot/simple-golang/database"
	"rendipilot/simple-golang/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Load environment variables
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Initialize database connection ONCE
	_, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	defer database.GetDB().Close(context.Background())

	config := cors.New(cors.Config{
		AllowOrigins:   "*",
		AllowMethods:   "GET,POST,HEAD,PUT,DELETE,PATCH", // Allow all methods
		AllowHeaders:   "",
		AllowCredentials: false,
	})

	app.Use(config)

	// Define endpoints
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello Rendy")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Test")
	})

	app.Post("/adduser", handlers.AddUser)

	app.Get("/users", handlers.GetUsers)

	// Start the server
	app.Listen(":3000")
}
