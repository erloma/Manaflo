package models

import (
	"time"
	"gorm.io/gorm"
)

type Priority string
type TaskStatus string
type ReviewStatus string
type Commentable string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"

    Unassigned TaskStatus = "unassigned"
    Todo TaskStatus = "todo"
    InProgress TaskStatus = "in_progress"
    Done TaskStatus = "done"  

    Pending ReviewStatus = "pending"
    Approved ReviewStatus = "approved"
    Rejected ReviewStatus = "rejected"

    CommTask Commentable = "task"
    CommAttachment Commentable = "attachment"
)

type Task struct {
    gorm.Model
    Project         uint         `json:"project"`
    Title           string       `gorm:"not null" json:"title"`
    Description     string       `json:"description"`
    Deadline        time.Time    `json:"deadline"`
    Priority        Priority     `gorm:"type:enum('low', 'medium', 'high')" json:"priority"`
    Weight          int          `json:"weight"`
    Status          TaskStatus   `gorm:"type:enum('unassigned', 'todo', 'in_progress', 'done')" json:"status"`
    AssignedUsers   []uint       `json:"assigned_users"`
    CreatedBy       uint         `json:"created_by"` 
    Attachments     []Attachment `json:"attachments"`
    Comments        []Comment    `json:"comments"`
}

type TaskReview struct {
    gorm.Model
    Task         Task            `gorm:"not null" json:"taskID"`
    ReviewerID   uint            `gorm:"not null" json:"reviewerID"`
    ReviewStatus ReviewStatus    `gorm:"type:ENUM('pending', 'approved', 'rejected');default:'pending'" json:"status"`
    Comment      string          `gorm:"type:text" json:"comment"`
    CommentsList []Comment       `json:"comments_list"`
}

type Comment struct {
    gorm.Model
    User            uint         `json:"user"`      
    CommentText     string       `gorm:"text" json:"comment_text"`
    CommentableID   uint         `json:"commentable_id"` 
    CommentableType Commentable  `gorm:"type:ENUM('task', 'attachment')" json:"commentable_type"`
}

type Attachment struct {
    gorm.Model
    FilePath   string    `json:"file_path"`  
    FileName   string    `json:"file_name"`
    FileType   string    `json:"file_type"`
    UploadedBy uint      `json:"uploaded_by"` 
    TaskID    uint       `json:"task_id"`
}
