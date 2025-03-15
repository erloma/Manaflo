package models

import (
	"time"
	"gorm.io/gorm"
)


type Task struct {
    gorm.Model
    Project         uint         `json:"project" validate:"required"`
    Title           string       `json:"title" validate:"required"`
    Description     string       `json:"description" validate:"required"`
    Deadline        *time.Time   `json:"deadline"`
    Priority        string       `gorm:"type:varchar(255);default:'medium'" json:"priority"`
    Weight          *int         `json:"weight"`
    Status          string       `gorm:"type:varchar(255);default:'unassigned'" json:"status"`
    CreatedBy       uint         `json:"created_by" validate:"required"` 
    AssignedUsers   []uint       `gorm:"type:json" json:"assigned_users"`
    Attachments     []Attachment `gorm:"type:json" json:"attachments"`
    Comments        []Comment    `gorm:"type:json" json:"comments"`

}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
    if t.Attachments == nil {
		t.Attachments = []Attachment{}
	}
    if t.AssignedUsers == nil {
		t.AssignedUsers = []uint{}
	}
    if t.Comments == nil {
        t.Comments = []Comment{}
    }
	return
}

type TaskReview struct {
    gorm.Model
    TaskID       uint            `gorm:"not null" json:"taskID"`
    ReviewerID   uint            `gorm:"not null" json:"reviewerID"`
    ReviewStatus string          `gorm:"type:varchar(255);default:'pending'" json:"status"`
    Comment      string          `gorm:"type:text" json:"comment"`
    CommentsList []Comment       `gorm:"type:json" json:"comments_list"`
}

type Comment struct {
    gorm.Model
    User            uint         `json:"user"`      
    CommentText     string       `gorm:"text" json:"comment_text"`
    CommentableID   uint         `json:"commentable_id"` 
    CommentableType string       `gorm:"type:varchar(255)" json:"commentable_type"`
}

type Attachment struct {
    gorm.Model
    FilePath   string    `json:"file_path"`  
    FileName   string    `json:"file_name"`
    FileType   string    `json:"file_type"`
    UploadedBy uint      `json:"uploaded_by"` 
    TaskID     uint      `json:"task_id"`
} 

// TODO: Fix enums for TaskStatus, TaskPriority, ReviewStatus and CommentableType
