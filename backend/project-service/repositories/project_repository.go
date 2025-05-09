package repositories

import (
	"github.com/erloma/manaflo/backend/project-service/config"
	"github.com/erloma/manaflo/backend/project-service/models"
	"gorm.io/gorm"
)

func CreateProject(project *models.Project) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	return db.Create(project).Error
}

func GetAllProjects() ([]models.Project, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var projects []models.Project
	if err := db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectByID(id uint) (*models.Project, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var project models.Project
	if err := db.First(&project, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func GetProjectMembers(id uint) ([]models.ProjectMember, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var projectMembers []models.ProjectMember

	err = db.Where("project_id = ?", id).Find(&projectMembers).Error

	if err != nil {
		return nil, err
	}
	return projectMembers, err
}
