package models

import "gorm.io/gorm"

type Conversation struct {
    ProjectID  uint          `json:"project_id"` 
    Name       string        `json:"name"`                 
    Messages   []Message     `json:"messages"`             
}

type Message struct {
    gorm.Model
    Conversation   uint      `json:"conversation_id"` 
    SenderID       uint      `json:"sender_id"`      
    MessageText    string    `json:"message_text"`
}