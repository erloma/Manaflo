package main

import (
	"fmt"
	"log"

	"github.com/erloma/manaflo/backend/project-service/config"
	"github.com/erloma/manaflo/backend/project-service/models"
	"github.com/erloma/manaflo/backend/project-service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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
    db.AutoMigrate(&models.Project{})

    // Setup routes
    routes.SetupRoutes(app)

	port := ":8082"

    // Start server
    log.Fatal(app.Listen(port))

	fmt.Println("Server is running on port", port)
}
