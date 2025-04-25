package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;size:16"         json:"firstName"`
	LastName  string `gorm:"not null;size:25"         json:"lastName"`
	Email     string `gorm:"not null;size:100;unique" json:"email"`
	Password  string `gorm:"not null;size:255"        json:"password"`
}

type LoginRequest struct {
	Email    string `gorm:"not null;size:100;unique" json:"email"`
	Password string `gorm:"not null;size:255"        json:"password"`
}

type UserUpdateRequest struct {
	FirstName *string `json:"firstName,omitempty"` //if not provided omit it. 
	LastName *string `json:"lastName,omitempty"` 
	Email *string `json:"email,omitempty"` 
	OldPassword *string `json:"oldPassword,omitempty"` 
	NewPasswordFirst *string `json:"newPasswordFirst,omitempty"` 
	NewPasswordSecond *string `json:"newPasswordSecond,omitempty"` 
}
