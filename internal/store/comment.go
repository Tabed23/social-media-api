package store

import (
	"github.com/tabed23/social-media-api/internal/repository"
	"gorm.io/gorm"
)

type CommentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) repository.CommentRepository {
	return &CommentStore{db: db}
}
