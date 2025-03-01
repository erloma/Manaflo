package models

import (
	"time"
	"gorm.io/gorm"
)


type Project struct {
    gorm.Model
    Name        string       `gorm:"not null;size:255" json:"name"`
    Description string       `gorm:"type:text"         json:"description"`
    CreatedBy   int          `gorm:"not null"          json:"created_by"`
    Attachments []Attachment `json:"attachments"`
}

type Role struct {
    gorm.Model
    Project     Project   `gorm:"not null"           json:"project"`           
    Name        string    `gorm:"not null;size:100"  json:"name"`  
    Description string    `gorm:"type:text"          json:"description"`
}

type ProjectMember struct {
    gorm.Model
    Project Project   `gorm:"not null"                  json:"project"`
    User    uint      `gorm:"not null"                  json:"user"`
    Role    Role      `gorm:"not null;default:'Member"  json:"role"`
}

type ProjectMilestone struct {
    ProjectID   Project   `json:"project"`
    Name        string    `json:"name"` 
    Description string    `json:"description"`
    DueDate     time.Time `json:"due_date"`
    Tasks       []uint    `json:"tasks"`
}

type Attachment struct {
    gorm.Model
    FilePath   string    `json:"file_path"`  
    FileName   string    `json:"file_name"`
    FileType   string    `json:"file_type"`
    UploadedBy uint      `json:"uploaded_by"` 
    ProjectID    uint    `json:"project_id"`
}

// Permissions for differing roles?