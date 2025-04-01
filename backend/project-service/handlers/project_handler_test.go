package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"github.com/erloma/manaflo/backend/project-service/config"
	"github.com/erloma/manaflo/backend/project-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func TestCreateProject(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(".env file not loaded")
	}

	db, err := config.GetDB()
	if err != nil {
		t.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&models.Project{})

	app := fiber.New()
	app.Post("/projects", CreateProject)
	project := models.Project{Name: "Test", Description: "Test project", CreatedBy: 1, Attachments: []models.Attachment{}}
	body, _ := json.Marshal((project))

	req := httptest.NewRequest("POST", "/projects", bytes.NewReader(body))
	req.Header.Set("Content-type", "application/json")

	// Send request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	// Check status
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

