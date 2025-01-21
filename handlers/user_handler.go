package handlers

import (
	"rendipilot/simple-golang/data"
	"rendipilot/simple-golang/models"

	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "error": err.Error(), })
    }

	if err := data.CreateUserDatabase(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User added successfully"})
}