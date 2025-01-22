package handlers

import (
	"log"
	"rendipilot/simple-golang/data"
	"rendipilot/simple-golang/models"

	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error convert json data ": err.Error(), })
    }

	if err := data.CreateUserDatabase(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error add data user to database ": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User added successfully"})
}

func GetUsers(c *fiber.Ctx) error {
	users, err := data.GetUsersData()
	if err != nil {
		// If an error occurs, log it and return a 500 internal server error response
		log.Println("Error fetching users:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve users",
		})
	}

	// Return the users data as JSON
	return c.Status(fiber.StatusOK).JSON(users)
}