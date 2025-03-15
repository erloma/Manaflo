package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/project-service/handlers"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/api/projects", handlers.CreateProject)
    app.Get("/api/projects", handlers.GetProjects)
}