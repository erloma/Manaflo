package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gababool/proman/backend/user-service/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/ping", handlers.PingHandler)

    app.Get("/api/users", handlers.GetUsers)
    app.Post("/api/users", handlers.CreateUser)

}