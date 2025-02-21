// main.go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gababool/proman/backend/user-service/internal/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/ping", handlers.PingHandler)

	port := ":8080"
	fmt.Println("Server is running on port", port)

	log.Fatal(app.Listen(port))
}
