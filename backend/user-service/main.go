package main

import (
	"fmt"
	"github.com/erloma/manaflo/backend/user-service/models"
	"log"

	"github.com/erloma/manaflo/backend/user-service/config"
	"github.com/erloma/manaflo/backend/user-service/handlers"
	"github.com/erloma/manaflo/backend/user-service/routes"
	"github.com/erloma/manaflo/backend/user-service/services"
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

	// Initialize app and middleware
	app := fiber.New()

	// CORS middleware setup
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
	db.AutoMigrate(&models.User{})

	// Initialize services and handlers
	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes and pass the handler
	routes.SetupRoutes(app, userHandler)

	// Start server
	port := ":8080"
	log.Fatal(app.Listen(port))

	fmt.Println("Server is running on port", port)
}
