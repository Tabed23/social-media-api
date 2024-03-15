package repository

import (
	"context"

	"github.com/tabed23/social-media-api/graph/models"
)

type UserRepository interface {
	UserCreate(context.Context, models.NewUserInput) (*models.User, error)
	GetUsers(context.Context) ([]*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	DeleteUser(context.Context, string) (bool, error)
	UpdateUser(context.Context, string, models.UpdateUserInput) (*models.User, error)
}

type CommentRepository interface{}
type PostRepository interface {
	CreaterPost(context.Context, string, models.NewPostInput) (*models.Post, error)
	GetPosts(context.Context) ([]*models.Post, error)
	GetPost(context.Context, string) (*models.Post, error)
	DeletePost(context.Context, string) (bool, error)
	UpdatePost(context.Context, string, models.UpdatePostInput) (*models.Post, error)
}

type LikeRepository interface{}
