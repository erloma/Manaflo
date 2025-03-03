package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/task-service/config"
    "github.com/erloma/manaflo/backend/task-service/models"
)


func CreateTask(c *fiber.Ctx) error {
    db, err := config.GetDB();
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
    }

    task := new(models.Task)

    if err := c.BodyParser(task); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Unable to parse JSON"})
    }

    if err := db.Create(task).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Unable create task"})
    }

    return c.JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
    db, err := config.GetDB();
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
    }

    tasks := new([]models.Task)

    if err := db.Find(tasks).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch tasks"})
    }

    return c.JSON(tasks)
}