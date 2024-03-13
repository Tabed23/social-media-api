package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tabed23/social-media-api/graph/models"
	"github.com/tabed23/social-media-api/internal/repository"
	"github.com/tabed23/social-media-api/internal/utils"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

// GetUsers implements repository.UserRepository.
func (u *UserStore) GetUsers(ctx context.Context) ([]*models.User, error) {
	usrs := []*models.User{}
	tx := u.db.Table("users").WithContext(ctx).Find(&usrs).Error
	if tx != nil {
		return nil, tx
	}
	return usrs, nil
}

// UserCreate implements repository.UserRepository.
func (u *UserStore) UserCreate(ctx context.Context, nuser models.NewUserInput) (*models.User, error) {
	pass := utils.HashPassword(nuser.Password)
	usr := models.User{
		ID:             uuid.New().String(),
		Username:       nuser.Username,
		Name:           nuser.Name,
		Email:          nuser.Email,
		ProfilePicture: nuser.ProfilePicture,
		Password:       pass,
		Bio:            nuser.Bio,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	tx := u.db.Table("users").Create(&usr).Error
	if tx != nil {
		return nil, tx
	}
	return &usr, nil
}

func NewUserStore(db *gorm.DB) repository.UserRepository {
	return &UserStore{db: db}
}
