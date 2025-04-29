package routes

import (
	"github.com/erloma/manaflo/backend/user-service/handlers"
	"github.com/erloma/manaflo/backend/user-service/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {

	// Non-protected routes
	app.Get("/ping", userHandler.PingHandler)
	app.Get("/api/users", userHandler.GetUsers)
	app.Post("/api/users", userHandler.CreateUser)
	app.Post("/api/login", userHandler.LoginUser)
	app.Patch("/api/users/:id", userHandler.UpdateUser)
	app.Get("/api/users/by-ids", userHandler.GetUsersByIds)

	// Protected routes (below this point)
	api := app.Group("/api", middleware.AuthRequired())
	api.Get("/user/profile", userHandler.GetUserProfile)
}
