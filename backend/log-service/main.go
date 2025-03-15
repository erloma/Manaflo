package main

import (
    "log"
	"fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/erloma/manaflo/backend/log-service/config"
    "github.com/erloma/manaflo/backend/log-service/models"
    "github.com/erloma/manaflo/backend/log-service/routes"
	"github.com/joho/godotenv"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system environment variables")
	}

    app := fiber.New()

    app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

    // Connect to database
    db, err := config.GetDB()
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto migrate the schema
    db.AutoMigrate(&models.Log{})

    // Setup routes
    routes.SetupRoutes(app)

	port := ":8084"

    // Start server
    log.Fatal(app.Listen(port))

	fmt.Println("Server is running on port", port)
}
