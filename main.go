package main

import (
	"log"
	"rendipilot/simple-golang/database"
	"rendipilot/simple-golang/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// connection to postgresql
	db, err := database.ConnectDatabase()

	if err != nil { 
		log.Fatal("Failed to connect to PostgreSQL:", err) 
	} 

	defer db.Close()

	// endpoint

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("hello rendy")
	})

	app.Post("/adduser", handlers.AddUser)

	// listen to port 3000 machine
	app.Listen(":3000")
}