package handlers

import (
	. "aquarium/database"
	. "aquarium/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(c fiber.Ctx) error {
	var users []User
	result := DB.Find(&users)
	fmt.Println(result)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}
	return c.JSON(users)
}

func CreateUser(c fiber.Ctx) error {
	user := new(User)
	if err := json.Unmarshal(c.Body(), user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// if err := c.Bind().Body(user); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	// }
	fmt.Println(user)
	result := DB.Create(user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).SendString(result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
