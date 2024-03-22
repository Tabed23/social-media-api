package store

import (
	"context"

	"github.com/tabed23/social-media-api/graph/models"
	"github.com/tabed23/social-media-api/internal/repository"
	"gorm.io/gorm"
)

type LikeStore struct {
	db *gorm.DB
}

func NewLikeStore(db *gorm.DB) repository.LikeRepository {
	return &LikeStore{db: db}
}

// ClickLike implements repository.LikeRepository.
func (l *LikeStore) ClickLike(context.Context, string, string, string) (*models.Like, error) {
	panic("unimplemented")
}

// ClickUnlike implements repository.LikeRepository.
func (l *LikeStore) ClickUnlike(context.Context, string) (bool, error) {
	panic("unimplemented")
}
