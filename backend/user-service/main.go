package main

import (
    "log"
	"fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gababool/proman/backend/user-service/config"
    "github.com/gababool/proman/backend/user-service/models"
    "github.com/gababool/proman/backend/user-service/routes"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system environment variables")
	}

    app := fiber.New()

    // Connect to database
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto migrate the schema
    db.AutoMigrate(&models.User{})

    // Setup routes
    routes.SetupRoutes(app)

	port := ":8080"

    // Start server
    log.Fatal(app.Listen(port))

	fmt.Println("Server is running on port", port)
}
