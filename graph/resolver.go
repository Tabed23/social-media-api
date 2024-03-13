package graph

import "github.com/tabed23/social-media-api/internal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repository.UserRepository
	repository.CommentRepository
	repository.LikeRepository
	repository.PostRepository
}
