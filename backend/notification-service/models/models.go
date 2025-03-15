package models

import (
	"gorm.io/gorm"
)

type Notification struct {
    gorm.Model
    User           uint        `json:"user"`           
    NotificationType string    `json:"notification_type"` // Add enum for this
    Message          string    `json:"message"`
    IsRead           bool      `json:"is_read"`
    RelatedProjectID uint      `json:"related_project_id,omitempty"`
    RelatedTaskID    uint      `json:"related_task_id,omitempty"`
}