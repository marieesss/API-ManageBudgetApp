package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marieesss/API-ManageBudgetApp/database"
	"github.com/marieesss/API-ManageBudgetApp/models"
	"github.com/marieesss/API-ManageBudgetApp/utils"
)

func CreateUser(c *fiber.Ctx) error {
	// instanciate new user
	user := new(models.User)

	// add the body values to user instance
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user.Password = hashedPassword

	// Create user
	database.Db.Db.Create(&user)

	return c.Status(200).JSON(user)

}
