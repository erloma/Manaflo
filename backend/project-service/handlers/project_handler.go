package handlers

import (
	"github.com/erloma/manaflo/backend/project-service/models"
	"github.com/erloma/manaflo/backend/project-service/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProjectHandler struct {
	service *services.ProjectService
}

func NewProjectHandler(svc *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: svc}
}

func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to parse JSON"})
	}

	if project.Name == "" || len(project.Name) > 255 || project.CreatedBy <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid fields in project entered"})
	}

	validTypes := map[string]bool{"Work project": true, "School project": true, "Personal project": true, "Other": true}
	if !validTypes[project.Type] {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid project type"})
	}

	// TODO check if project.CreatedBy maps to real user

	if err := h.service.CreateProject(project); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to create project"})
	}

	return c.Status(200).JSON(project)
}

func (h *ProjectHandler) GetProjects(c *fiber.Ctx) error {
	projects, err := h.service.GetAllProjects()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch projects"})
	}
	return c.Status(200).JSON(projects)

}

func (h *ProjectHandler) GetProjectUsers(c *fiber.Ctx) error {
	projectIDParam := c.Params("project_id")
	if projectIDParam == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID is missing in request"})
	}


	projectIDUint64, err := strconv.ParseUint(projectIDParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid project ID"})
	}
	projectID := uint(projectIDUint64)

	users, err := h.service.GetProjectUsers(projectID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch users of project"})
	}
	return c.Status(200).JSON(users)
}

// recieve body send to service and ensure correct json reponse
/*
func getUserProjects(c *fiber.Ctx, userID int) error {
	db, err := config.GetDB()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
	}

	// Implement this shit
}

func getProjectByID(){

}

*/
