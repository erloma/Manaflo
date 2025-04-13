package repositories

import (
	"errors"
	"github.com/erloma/manaflo/backend/user-service/config"
	"github.com/erloma/manaflo/backend/user-service/models"
	"gorm.io/gorm"
)

func CreateUser(user models.User) (models.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return models.User{}, err
	}
	result := db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUserByID(id string) (models.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}
		return models.User{}, result.Error
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
