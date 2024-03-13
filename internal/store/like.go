package store

import (
	"github.com/tabed23/social-media-api/internal/repository"
	"gorm.io/gorm"
)

type LikeStore struct {
	db *gorm.DB
}

func NewLikeStore(db *gorm.DB) repository.LikeRepository {
	return &LikeStore{db: db}
}
