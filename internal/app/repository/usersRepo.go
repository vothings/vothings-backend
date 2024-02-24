package repository

import (
	"fmt"
	"vothings/internal/app/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// Insert implements service.UserRepo.
func (u *userRepo) Insert(user model.Users) error {
	user.ID = uuid.New()
	result := u.db.Create(&user)

	if result.Error != nil {
		return fmt.Errorf("error in insert Repository: %s", result.Error)
	}
	return nil
}

// GetByUsername implements service.UserRepo.
func (u *userRepo) GetByUsername(username string) (model.Users, error) {
	var user model.Users
	result := u.db.Where("user_name = ?", username).First(&user)

	if result.Error != nil {
		return model.Users{}, fmt.Errorf("error in GetByUsername Repository: %s", result.Error)
	}
	return user, nil
}

func NewRepository(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}
