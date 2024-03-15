package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             string    `gorm:"type:varchar(255);primaryKey"`
	Username       string    `gorm:"type:varchar(255);unique;not null"`
	Email          string    `gorm:"type:varchar(255);unique;not null"`
	Name           string    `gorm:"type:varchar(100)"`
	Bio            string    `gorm:"type:text"`
	ProfilePicture string    `gorm:"type:varchar(255)"`
	Password       string    `gorm:"type:varchar(255)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	// Relationships
	Posts        []Post    `gorm:"foreignKey:AuthorID"`
	Comments     []Comment `gorm:"foreignKey:AuthorID"`
	Likes        []Like    `gorm:"foreignKey:UserID"`
	Friends      []User    `gorm:"many2many:user_friends"`
	BlockedUsers []User    `gorm:"many2many:user_blocked_users"`
	CloseFriends []User    `gorm:"many2many:user_close_friends"`
}
