package services

import (
	"go-jwt/models"
)

type IUserService interface {
	Register(u *models.User) (user *models.User, err error)
	Login(username, password string) (token string, err error)
}

type IProductService interface {
	CreateProduct(p *models.Product, userID uint) (product *models.Product, err error)
	ReadAllProduct() (product *[]models.Product, err error)
	ReadProduct(productID int) (product *models.Product, err error)
	UpdateProduct(p *models.Product, productID int) (product *models.Product, err error)
	DeleteProduct(productID int) (err error)
}
