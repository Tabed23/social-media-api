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

func NewUserStore(db *gorm.DB) repository.UserRepository {
	return &UserStore{db: db}
}

// GetUsers implements repository.UserRepository.
func (u *UserStore) GetUsers(ctx context.Context) ([]*models.User, error) {
	usrs := []*models.User{}
	tx := u.db.Table("users").WithContext(ctx).Find(&usrs).Error
	if tx != nil {
		return nil, tx
	}
	for _, user := range usrs {
		if user == nil {
			continue
		}
		tx := u.db.Table("posts").WithContext(ctx).Where("author_id = ?", user.ID).Find(&user.Posts)
		if tx.Error != nil {
			return nil, tx.Error
		}
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
	tx := u.db.Table("users").WithContext(ctx).Create(&usr).Error
	if tx != nil {
		return nil, tx
	}
	return &usr, nil
}

// FindByEmail implements repository.UserRepository.
func (u *UserStore) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	usr := models.User{}
	if err := u.db.Table("users").WithContext(ctx).Where("email = ?", email).Find(&usr).Error; err != nil {
		return nil, err
	}
	tx := u.db.Table("posts").WithContext(ctx).Where("author_id = ?", usr.ID).Find(&usr.Posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &usr, nil
}

// FindByUsername implements repository.UserRepository.
func (u *UserStore) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	usr := models.User{}
	if err := u.db.Table("users").WithContext(ctx).Where("username = ?", username).Find(&usr).Error; err != nil {
		return nil, err
	}
	tx := u.db.Table("posts").WithContext(ctx).Where("author_id = ?", usr.ID).Find(&usr.Posts)
	if tx != nil {
		return nil, tx.Error
	}

	return &usr, nil
}

// DeleteUser implements repository.UserRepository.
func (u *UserStore) DeleteUser(ctx context.Context, email string) (bool, error) {
	if _, err := u.FindByEmail(ctx, email); err != nil {
		return false, err
	}
	if err := u.db.Table("users").WithContext(ctx).Delete(&models.User{}, "email=?", email).Error; err != nil {
		return false, err
	}
	return true, nil
}

// UpdateUser implements repository.UserRepository.
func (u *UserStore) UpdateUser(ctx context.Context, email string, update models.UpdateUserInput) (*models.User, error) {
	found, err := u.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if update.Password != nil {
		hash := utils.HashPassword(*update.Password)
		update.Password = &hash
	}
	updatedTime := time.Now().UTC()
	update.UpdatedAt = &updatedTime

	err = u.db.Table("users").WithContext(ctx).Where("email=?", found.Email).Updates(&update).Error
	if err != nil {
		return nil, err
	}

	return u.FindByEmail(ctx, email)

}
