package repository

import (
	"context"

	"github.com/tabed23/social-media-api/graph/models"
)

type UserRepository interface{
	UserCreate(context.Context, models.NewUserInput) (*models.User, error)
	GetUsers(context.Context)([]*models.User, error)
}

type CommentRepository interface{}
type PostRepository interface{}
type LikeRepository interface{}
