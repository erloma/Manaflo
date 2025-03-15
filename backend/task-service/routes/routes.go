package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/task-service/handlers"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/api/tasks", handlers.CreateTask)
    app.Get("/api/tasks", handlers.GetTasks)
}