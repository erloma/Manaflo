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

func (s *UserService) updateUser(update models.UserUpdateRequest) error {
	
	user, err := repositories.GetUserByID(userID) //check if user exists
	
	if err != nil { //if error we exit with the error
		return err
	}

	if update.Email != nil {
		if utils.CheckValidEmail(update.Email) {
			return errors.New("email is of invalid format")
		}
		user.Email = *update.Email; 
	}

	if update.firstName != nil {
		user.firstName = *update.firstName; 
	}

	if update.lastName != nil {
		user.lastName = *update.lastName; 
	}
	
	if update.newPasswordFirst != nil {
		if update.OldPassword == nil {
			return errors.New("Old password not provided")
		} 
		if update.newPasswordFirst != update.newPasswordSecond {
			return errors.New("New password does not match")
		}
		if !utils.CheckPasswordHash(update.OldPassword, user.Password) {
			return errors.New("Old password is incorrect")
		}
		if !utils.CheckPasswordValidity(update.Password) {
			return errors.New("password must contain a special character and be at least 8 characters")
		}
		hashedPassword, err := utils.HashPassword(user.Password)
		
		if err != nil {
			return err
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
