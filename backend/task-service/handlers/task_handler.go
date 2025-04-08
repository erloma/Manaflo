package handlers

import (
	"github.com/erloma/manaflo/backend/task-service/config"
	"github.com/erloma/manaflo/backend/task-service/models"
	"github.com/erloma/manaflo/backend/task-service/services"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	tasks, err := h.taskService.GetTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	db, err := config.GetDB()
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
	db, err := config.GetDB()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
	}

	tasks := new([]models.Task)

	if err := db.Find(tasks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch tasks"})
	}

	return c.JSON(tasks)
}

