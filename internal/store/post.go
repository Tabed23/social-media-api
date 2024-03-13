package store

import (
	"github.com/tabed23/social-media-api/internal/repository"
	"gorm.io/gorm"
)

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) repository.PostRepository {
	return &PostStore{db: db}
}
