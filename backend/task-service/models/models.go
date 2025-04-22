package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Project     uint       `json:"project" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Deadline    *time.Time `json:"deadline"`
	Priority    string     `gorm:"type:varchar(255);default:'medium'" json:"priority"`
	Weight      *int       `json:"weight"`
	Status      string     `gorm:"type:varchar(255);default:'unassigned'" json:"status"`
	CreatedBy   uint       `json:"created_by" validate:"required"`
	AssignedTo  uint       `json:"assigned_to"`
}
