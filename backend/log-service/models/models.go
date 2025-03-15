package models

import "gorm.io/gorm"

type Log struct {
    gorm.Model
    ServiceName string    `json:"service_name"` 
    LogLevel    string    `json:"log_level"`   // Add enum here
    Message     string    `json:"message"`
}