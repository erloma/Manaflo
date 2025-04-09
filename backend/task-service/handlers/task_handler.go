package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/task-service/config"
    "github.com/erloma/manaflo/backend/task-service/models"
    "github.com/erloma/manaflo/backend/task-service/services"
)

type TaskHandler struct {
    taskService *services.TaskService
}

func new TaskHandler(taskService *services.TaskService) *TaskHandler {
    return &TaskHandler{taskServier: taskService}
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
    tasks, err := h.taskService.GetTasks()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(tasks)
}



func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
    var task models.Task

    if err := c.BodyParser(task); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Unable to parse JSON"})
    }

    _, err := h.taskService.CreateTask(task)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Task created successfully"})
}
