package repository

import (
	"go-jwt/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(db *gorm.DB, user *models.User) (err error)
	Login(db *gorm.DB, username string) (user *models.User, err error)
}

type IProductRepository interface {
	CreateProduct(db *gorm.DB, p *models.Product) (product *models.Product, err error)
	ReadAllProduct(db *gorm.DB) (product *[]models.Product, err error)
	ReadProduct(db *gorm.DB, productID int) (product *models.Product, err error)
	UpdateProduct(db *gorm.DB, p *models.Product, productID int) (product *models.Product, err error)
	DeleteProduct(db *gorm.DB, productID int) (err error)
}
