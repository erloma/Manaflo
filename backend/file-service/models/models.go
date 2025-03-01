package models

import "gorm.io/gorm"

type File struct {
    gorm.Model
    FilePath   string    `json:"file_path"`
    FileName   string    `json:"file_name"`
    FileType   string    `json:"file_type"`
    FileSize   int64     `json:"file_size"` 
    UploadedBy uint      `json:"uploaded_by"`
}