package routes

import (
	"github.com/erloma/manaflo/backend/project-service/handlers"
	"github.com/erloma/manaflo/backend/project-service/services"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	svc := services.NewProjectService()
	h := handlers.NewProjectHandler(svc)
	app.Post("/api/projects", h.CreateProject)
	app.Get("/api/projects", h.GetProjects)
	api := app.Group("/api", middleware.AuthRequired())
	api.Get("/projects/users", h.GetProjectUsers)
}
