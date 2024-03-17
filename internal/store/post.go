package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tabed23/social-media-api/graph/models"
	"github.com/tabed23/social-media-api/internal/repository"
	"gorm.io/gorm"
)

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) repository.PostRepository {
	return &PostStore{db: db}
}

// CreaterPost implements repository.PostRepository.
func (p *PostStore) CreaterPost(ctx context.Context, username string, npost models.NewPostInput) (*models.Post, error) {
	usr := models.User{}
	if err := p.db.Table("users").WithContext(ctx).Where("username = ?", username).Find(&usr).Error; err != nil {
		return nil, err
	}
	post := models.Post{
		ID:        uuid.New().String(),
		Content:   npost.Content,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		AuthorID:  usr.ID,
		Author:    usr,
	}

	if err := p.db.Table("posts").WithContext(ctx).Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil

}

// GetPost implements repository.PostRepository.
func (p *PostStore) GetPost(ctx context.Context, postID string) (*models.Post, error) {
	post := &models.Post{}
	tx := p.db.Table("posts").WithContext(ctx).Where("id = ?", postID).Find(&post).Error
	if tx != nil {
		return nil, tx
	}
	tx = p.db.Table("comments").WithContext(ctx).Where("id = ?", post.AuthorID).Find(&post.Author).Error
	if tx != nil {
		return nil, tx
	}
	return post, nil
}

// GetPosts implements repository.PostRepository.
func (p *PostStore) GetPosts(ctx context.Context) ([]*models.Post, error) {
	posts := []*models.Post{}
	tx := p.db.Table("posts").WithContext(ctx).Find(&posts).Error
	if tx != nil {
		return nil, tx
	}
	for _, post := range posts {
		if post == nil {
			continue
		}
		tx = p.db.Table("users").WithContext(ctx).Where("id = ?", post.AuthorID).Find(&post.Author).Error
		if tx != nil {
			return nil, tx
		}

	}

	return posts, nil

}

// DeletePost implements repository.PostRepository.
func (p *PostStore) DeletePost(ctx context.Context, id string) (bool, error) {
	_, err := p.GetPost(ctx, id)
	if err != nil {
		return false, err
	}
	if err := p.db.Table("posts").WithContext(ctx).Delete(&models.Post{}, "id=?", id).Error; err != nil {
		return false, err
	}
	return true, nil

}

// UpdatePost implements repository.PostRepository.
func (p *PostStore) UpdatePost(ctx context.Context, id string, updated models.UpdatePostInput) (*models.Post, error) {
	_, err := p.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}
	updatedTime := time.Now().UTC()
	updated.UpdatedAt = &updatedTime
	err = p.db.Table("posts").WithContext(ctx).Where("id=?", id).Updates(&updated).Error
	if err != nil {
		return nil, err
	}
	return p.GetPost(ctx, id)
}
