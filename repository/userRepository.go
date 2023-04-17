package repository

import (
	"go-jwt/models"

	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository UserRepository) Login(db *gorm.DB, email string) (user *models.User, err error) {
	err = db.Debug().Where("email = ?", email).Take(&user).Error
	return user, err
}

func (r UserRepository) Register(db *gorm.DB, user *models.User) (err error) {
	err = db.Debug().Create(&user).Error
	return err
}
