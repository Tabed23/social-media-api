package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/tabed23/social-media-api/graph/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, newUser models.NewUserInput) (*models.User, error) {
	return r.UserRepository.UserCreate(ctx, newUser)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, email string, updateUser models.UpdateUserInput) (*models.User, error) {
	return r.UserRepository.UpdateUser(ctx, email, updateUser)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, email string) (bool, error) {
	return r.UserRepository.DeleteUser(ctx, email)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, username string, newPost models.NewPostInput) (*models.Post, error) {
	return r.PostRepository.CreaterPost(ctx, username, newPost)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, updatePost models.UpdatePostInput) (*models.Post, error) {
	panic(fmt.Errorf("not implemented: UpdatePost - updatePost"))
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	return r.PostRepository.DeletePost(ctx, id)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, username string, newComment models.NewCommentInput) (*models.Comment, error) {
	return r.CommentRepository.CreaterComment(ctx, username, newComment)
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, updateComment models.UpdateCommentInput) (*models.Comment, error) {
	return r.CommentRepository.UpdateComment(ctx, id, updateComment)
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	return r.CommentRepository.DeleteComment(ctx, id)
}

// CreateLike is the resolver for the createLike field.
func (r *mutationResolver) CreateLike(ctx context.Context, userID string, postID *string, commentID *string) (*models.Like, error) {
	panic(fmt.Errorf("not implemented: CreateLike - createLike"))
}

// Unlike is the resolver for the unlike field.
func (r *mutationResolver) Unlike(ctx context.Context, likeID string) (bool, error) {
	panic(fmt.Errorf("not implemented: Unlike - unlike"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, search *string, limit *int, offset *int) ([]*models.User, error) {
	return r.UserRepository.GetUsers(ctx)
}

// Userbyusername is the resolver for the userbyusername field.
func (r *queryResolver) Userbyusername(ctx context.Context, username string) (*models.User, error) {
	return r.UserRepository.FindByUsername(ctx, username)
}

// Userbyemail is the resolver for the userbyemail field.
func (r *queryResolver) Userbyemail(ctx context.Context, email string) (*models.User, error) {
	return r.UserRepository.FindByEmail(ctx, email)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*models.Post, error) {
	return r.PostRepository.GetPosts(ctx)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	return r.PostRepository.GetPost(ctx, id)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, username string, limit *int, offset *int) ([]*models.Comment, error) {
	return r.CommentRepository.GetComments(ctx, username)
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, usename string, id string) (*models.Comment, error) {
	return r.CommentRepository.GetComment(ctx, usename, id)
}

// CommentsByPost is the resolver for the commentsByPost field.
func (r *queryResolver) CommentsByPost(ctx context.Context, postID string) ([]*models.Comment, error) {
	return r.CommentRepository.GetCommentByPost(ctx, postID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
