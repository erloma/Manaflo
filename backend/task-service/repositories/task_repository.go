package repositories

import (
		"github.com/erloma/manaflo/backend/task-service/config"
		"github.com/erloma/manaflo/backend/task-service/models"
)

func CreateTask(task models.Task) (models.Task, error) {
	db, err := config.GetDB()
	if err != nil {
		return models.Task{}, err
	}
	result := db.Create(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}
// Refactor later to only query tasks in project ID
func GetTasks() ([]models.Task, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}	