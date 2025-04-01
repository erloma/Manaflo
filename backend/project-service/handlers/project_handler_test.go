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
	"github.com/stretchr/testify/assert"
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

func TestGetProjects(t *testing.T) {
	// Load env
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Couldnt load env file")
	}

	// Connect to and migrate DB
	db, err := config.GetDB()
	if err != nil {
		t.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&models.Project{})

	// Create a test project to get from DB
	testProject := models.Project{Name: "ProjectTestName", Description: "Description", CreatedBy: 2}
	db.Create(&testProject)

	// Setup app and route
	app := fiber.New()
	app.Get("/projects", GetProjects)

	// Simulate GET request
	req := httptest.NewRequest("GET", "/projects", nil)
	req.Header.Set("Content-type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Decode resp body
	var projects []models.Project
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		t.Fatalf("error decoding body from JSON")
	}

	// Assert at least one project is returned
	assert.GreaterOrEqual(t, len(projects), 1)

	foundBool := false
	// Check that a project with the test name exists
	for _, p := range projects {
		if p.Name == testProject.Name && p.CreatedBy == testProject.CreatedBy && p.Description == testProject.Description {
			foundBool = true
			break
		}
	}
	assert.True(t, foundBool, "Test project not found in the response")

	// Clear the DB after the test
	db.Where("name = ?", "ProjectTestName").Delete(&models.Project{})

}
