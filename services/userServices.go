package services

import (
	"go-jwt/helpers"
	"go-jwt/models"
	repo "go-jwt/repository"

	"gorm.io/gorm"
)

type UserService struct {
	UserRepository repo.IUserRepository
	DB             *gorm.DB
}

func NewUserService(repository repo.IUserRepository, db *gorm.DB) *UserService {
	return &UserService{
		UserRepository: repository,
		DB:             db,
	}
}

func (service UserService) Login(email, password string) (token string, err error) {
	user, err := service.UserRepository.Login(service.DB, email)
	if err != nil || !helpers.ComparePass([]byte(user.Password), []byte(password)) {
		return "", err
	}
	token = helpers.GenerateToken(user.ID, user.Email, user.Role)

	return token, nil
}

func (service UserService) Register(u *models.User) (user *models.User, err error) {
	err = service.UserRepository.Register(service.DB, u)
	if err != nil {
		return nil, err
	}
	user = u
	return user, nil
}
