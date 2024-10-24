package controllers_v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rochi88/goapi/app/helpers"
	"github.com/rochi88/goapi/app/models"
)

func GetUser(c *fiber.Ctx) error {
	db := helpers.DB
	id := c.Params("id")
	var user models.User
	db.Find(&user, id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "No user found with ID",
		})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := helpers.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	db.Create(&user)
	return c.JSON(user)
}
