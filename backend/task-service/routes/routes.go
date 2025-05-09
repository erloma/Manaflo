package routes

import (
	"github.com/erloma/manaflo/backend/task-service/handlers"
	"github.com/erloma/manaflo/backend/task-service/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, taskHandler *handlers.TaskHandler) {

	api := app.Group("/api", middleware.AuthRequired())
	api.Post("/api/tasks", taskHandler.CreateTask)
	api.Get("/api/tasks", taskHandler.GetTasks)

}

