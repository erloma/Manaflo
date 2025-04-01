package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/task-service/middleware"
    "github.com/erloma/manaflo/backend/task-service/handlers"
)

func SetupRoutes(app *fiber.App) {


    // Protected routes (below this point)
    api := app.Group("/api", middleware.AuthRequired())
    api.Post("/api/tasks", handlers.CreateTask)
    api.Get("/api/tasks", handlers.GetTasks)
}