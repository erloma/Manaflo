package services

import (
	"errors"
	"github.com/erloma/manaflo/backend/project-service/external"
	"github.com/erloma/manaflo/backend/project-service/models"
	"github.com/erloma/manaflo/backend/project-service/repositories"
)

type ProjectService struct{}

type ProjectMember struct{}

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

func (s *ProjectService) GetProjectUsers(id uint) ([]external.UserInfo, error) {
	projectMembers, err := repositories.GetProjectMembers(id)
	if err != nil {
		return nil, err
	}

	var users []uint
	for _, member := range projectMembers {
		users = append(users, member.User)
	}

	var members []external.UserInfo

	members, err = external.GetUserInfosByIDs(users)

	if err != nil {
		return nil, err
	}

	return members, nil
}
