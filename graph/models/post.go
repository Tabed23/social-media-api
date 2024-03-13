package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        string    `gorm:"type:varchar(255);primaryKey"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	AuthorID  string    `gorm:"type:varchar(255);not null"`

	// Relationships
	Author   User      `gorm:"foreignKey:AuthorID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
	Likes    []Like    `gorm:"foreignKey:PostID"`
}
