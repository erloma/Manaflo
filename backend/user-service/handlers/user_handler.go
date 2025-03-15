package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/user-service/config"
    "github.com/erloma/manaflo/backend/user-service/models"
)

func PingHandler(c *fiber.Ctx) error {
	return c.SendString("Server is running!")
}

func GetUsers(c *fiber.Ctx) error {
    db, err := config.GetDB()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
    }

    var users []models.User
    if err := db.Find(&users).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
    }
    return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
    db, err := config.GetDB()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
    }

    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := db.Create(&user).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
    }
    return c.JSON(user)
}
