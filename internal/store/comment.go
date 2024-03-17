package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tabed23/social-media-api/graph/models"
	"github.com/tabed23/social-media-api/internal/repository"
	"gorm.io/gorm"
)

type CommentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) repository.CommentRepository {
	return &CommentStore{db: db}
}

// CreaterPost implements repository.CommentRepository.
func (c *CommentStore) CreaterComment(ctx context.Context, username string, comment models.NewCommentInput) (*models.Comment, error) {
	usr := models.User{}
	if err := c.db.Table("users").WithContext(ctx).Where("username = ?", username).Find(&usr).Error; err != nil {
		return nil, err
	}
	post := models.Post{}
	if err := c.db.Table("posts").WithContext(ctx).Where("id=?", comment.PostID).Find(&post).Error; err != nil {
		return nil, err
	}
	cmt := models.Comment{
		ID:        uuid.New().String(),
		Content:   comment.Content,
		PostID:    comment.PostID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Post:      post,
		AuthorID:  usr.ID,
		Author:    usr,
	}
	if err := c.db.Table("comments").WithContext(ctx).Create(&cmt).Error; err != nil {
		return nil, err
	}

	return &cmt, nil
}

// DeleteComment implements repository.CommentRepository.
func (c *CommentStore) DeleteComment(ctx context.Context, cmtID string) (bool, error) {
	err := c.db.Table("comments").WithContext(ctx).Delete(&models.Comment{}, "id=?", cmtID).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetComment implements repository.CommentRepository.
func (c *CommentStore) GetComment(ctx context.Context, usrname string, postID string) (*models.Comment, error) {

	var user models.User
	if err := c.db.Table("users").WithContext(ctx).Where("username = ?", usrname).First(&user).Error; err != nil {
		return nil, err
	}
	var comment models.Comment

	if err := c.db.Table("comments").WithContext(ctx).Select("comments.*").
		Joins("join posts on posts.id = comments.post_id").
		Where("comments.author_id = ? AND comments.post_id = ?", user.ID, postID).
		First(&comment).Error; err != nil {
		return nil, err
	}
	if err := c.db.Table("users").WithContext(ctx).Where("id = ?", comment.AuthorID).Find(&comment.Author).Error; err != nil {
		return nil, err
	}
	if err := c.db.Table("posts").WithContext(ctx).Where("id=?", comment.PostID).Find(&comment.Post).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// GetCommentByPost implements repository.CommentRepository.
func (c *CommentStore) GetCommentByPost(ctx context.Context, postid string) ([]*models.Comment, error) {
	var post models.Post
	if err := c.db.Table("posts").WithContext(ctx).Where("post_id = ?", postid).Find(&post).Error; err != nil {
		return nil, err
	}
	var comments []*models.Comment
	err := c.db.Table("comments").WithContext(ctx).Where("post_id = ?", post.ID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for _, cmt := range comments {
		if cmt == nil {
			continue
		}
		if err := c.db.Table("users").WithContext(ctx).Where("id = ?", cmt.AuthorID).Find(&cmt.Author).Error; err != nil {
			return nil, err
		}
		if err := c.db.Table("posts").WithContext(ctx).Where("id=?", cmt.PostID).Find(&cmt.Post).Error; err != nil {
			return nil, err
		}
	}
	return comments, nil
}

// GetComments implements repository.CommentRepository.
func (c *CommentStore) GetComments(ctx context.Context, username string) ([]*models.Comment, error) {
	var comments []*models.Comment
	var user models.User
	if err := c.db.WithContext(ctx).Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	if err := c.db.WithContext(ctx).Table("comments").Where("author_id = ?", user.ID).Find(&comments).Error; err != nil {
		return nil, err
	}
	for _, cmt := range comments {
		if cmt == nil {
			continue
		}
		if err := c.db.Table("users").WithContext(ctx).Where("username = ?", username).Find(&cmt.Author).Error; err != nil {
			return nil, err
		}
		if err := c.db.Table("posts").WithContext(ctx).Where("id=?", cmt.PostID).Find(&cmt.Post).Error; err != nil {
			return nil, err
		}
	}

	return comments, nil

}

// UpdateComment implements repository.CommentRepository.
func (c *CommentStore) UpdateComment(ctx context.Context, id string, updated models.UpdateCommentInput) (*models.Comment, error) {

	updatedTime := time.Now().UTC()
	updated.UpdatedAt = &updatedTime
	err := c.db.Table("comments").WithContext(ctx).Where("id=?", id).Updates(&updated).Error
	if err != nil {
		return nil, err
	}
	var cmt models.Comment
	if err := c.db.WithContext(ctx).Table("comments").Where("id = ?", id).Find(&cmt).Error; err != nil {
		return nil, err
	}
	return &cmt, nil
}
