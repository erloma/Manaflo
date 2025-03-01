package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName string `gorm:"not null;size:16"         json:"firstName"`
	LastName  string `gorm:"not null;size:25"         json:"lastName"`
    Email     string `gorm:"not null;size:100;unique" json:"email"`
	Password  string `gorm:"not null;size:25"         json:"password"`
    
    // Profile image?
}