// main.go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gababool/proman/backend/user-service/internal/handlers"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Enable cors for frontend requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	app.Get("/ping", handlers.PingHandler)

	port := ":8080"
	fmt.Println("Server is running on port", port)

	log.Fatal(app.Listen(port))
}
