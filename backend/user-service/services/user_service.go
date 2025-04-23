package services

import (
	"errors"
	"github.com/erloma/manaflo/backend/user-service/models"
	"github.com/erloma/manaflo/backend/user-service/repositories"
	"github.com/erloma/manaflo/backend/user-service/utils"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}


func (s *UserService) GetUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func (s *UserService) UpdateUser(userID string, update models.UserUpdateRequest) (models.User, error) {
	
	user, err := repositories.GetUserByID(userID) //check if user exists
	
	if err != nil { //if error we exit with the error
		return models.User{}, err
	}

	if update.Email != nil {
		if utils.CheckValidEmail(*update.Email) {
			return models.User{}, errors.New("email is of invalid format")
		}
		user.Email = *update.Email; 
	}

	if update.FirstName != nil {
		user.FirstName = *update.FirstName; 
	}

	if update.LastName != nil {
		user.LastName = *update.LastName; 
	}
	
	if update.NewPasswordFirst != nil {
		if update.OldPassword == nil {
			return models.User{}, errors.New("Old password not provided")
		} 
		if *update.NewPasswordFirst != *update.NewPasswordSecond {
			return models.User{}, errors.New("New password does not match")
		}
		if !utils.CheckPasswordHash(*update.OldPassword, user.Password) {
			return models.User{}, errors.New("Old password is incorrect")
		}
		if !utils.CheckPasswordValidity(*update.NewPasswordFirst) {
			return models.User{}, errors.New("password must contain a special character and be at least 8 characters")
		}
		hashedPassword, err := utils.HashPassword(*update.NewPasswordFirst)
		
		if err != nil {
			return models.User{}, err
		}
		user.Password = hashedPassword
	}
	return repositories.UpdateUser(user)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	if !utils.CheckValidEmail(user.Email) {
		return models.User{}, errors.New("email is of invalid format")
	}
	if !utils.CheckPasswordValidity(user.Password) {
		return models.User{}, errors.New("password must contain a special character and be at least 8 characters")
	}
	existingUser, err := repositories.GetUserByEmail(user.Email)
	if err != nil {
		return models.User{}, err
	}
	if existingUser.ID != 0 {
		return models.User{}, errors.New("user with this email already exists")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = hashedPassword
	return repositories.CreateUser(user)
}

func (s *UserService) LoginUser(request models.LoginRequest) (models.User, error) {
	user, err := repositories.GetUserByEmail(request.Email)
	if err != nil {
		return models.User{}, err
	}
	if utils.CheckPasswordHash(request.Password, user.Password) {
		return user, nil
	}
	return models.User{}, errors.New("incorrect password")
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
