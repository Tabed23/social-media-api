package models

import (
	"time"

	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	ID        string    `gorm:"type:varchar(255);primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relationships
	UserID    string `gorm:"type:varchar(255);not null"`
	PostID    string `gorm:"type:varchar(255)"`
	CommentID string `gorm:"type:varchar(255)"`

	User    User    `gorm:"foreignKey:UserID"`
	Post    Post    `gorm:"foreignKey:PostID"`
	Comment Comment `gorm:"foreignKey:CommentID"`
}
