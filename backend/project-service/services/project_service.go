package services

import (
	"errors"
	"github.com/erloma/manaflo/backend/project-service/models"
	"github.com/erloma/manaflo/backend/project-service/repositories"
)

type ProjectService struct {}

type UserInfo struct {
	id uint
	FirstName string
	LastName string
}
func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) CreateProject(project *models.Project) error {
	if project.Name == "" {
		return errors.New("Name cannot be empty")
	}
	if len(project.Name) > 255 {
		return errors.New("Name cannot be longer than 255 chars")
	}
	if project.CreatedBy <= 0 {
		return errors.New("Invalid userID")
	}

	// TODO verify project.CreatedBy exists in the users table

	return repositories.CreateProject(project)
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	return repositories.GetAllProjects()
}

func (s *ProjectService) GetProjectByID(id uint) (*models.Project, error) {
	return repositories.GetProjectByID(id)
}
func (s *ProjectService) GetProjectUsers(id uint) ([]UserInfo, error) {
	
}
