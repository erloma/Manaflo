package handlers

import (
	"github.com/erloma/manaflo/backend/project-service/config"
	"github.com/erloma/manaflo/backend/project-service/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	db, err := config.GetDB()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
	}

	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to parse JSON"})
	}

    if project.Name == "" || len(project.Name) > 255 || project.CreatedBy <= 0 {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid fields in project entered"})
    }

    // TODO add check if createdBy ID maps to a valid user

	if err := db.Create(project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to create project"})
	}

	return c.JSON(project)
}

func GetProjects(c *fiber.Ctx) error {
	db, err := config.GetDB()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database connection failed"})
	}

	projects := new([]models.Project)

	if err := db.Find(projects).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Unable to fetch projects"})
	}

	return c.JSON(projects)

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