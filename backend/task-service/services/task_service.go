package services

import (
	"github.com/erloma/manaflo/backend/task-service/models"
	"github.com/erloma/manaflo/backend/task-service/repositories"
)

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return repositories.GetTasks()
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return repositories.CreateTask(task)
}

