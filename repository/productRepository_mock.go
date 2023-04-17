package repository

import (
	"go-jwt/models"

	"github.com/stretchr/testify/mock"

	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) ReadProduct(db *gorm.DB, productID int) (product *models.Product, err error) {
	arguments := repository.Mock.Called(productID)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	product = arguments.Get(0).(*models.Product)

	return product, nil
}

func (repository *ProductRepositoryMock) ReadAllProduct(db *gorm.DB) (product *[]models.Product, err error) {
	arguments := repository.Mock.Called()

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	product = arguments.Get(0).(*[]models.Product)

	return product, nil
}

func (repository *ProductRepositoryMock) CreateProduct(db *gorm.DB, p *models.Product) (product *models.Product, err error) {
	return nil, nil
}
func (repository *ProductRepositoryMock) UpdateProduct(db *gorm.DB, p *models.Product, productID int) (product *models.Product, err error) {
	return nil, nil
}
func (repository *ProductRepositoryMock) DeleteProduct(db *gorm.DB, productID int) (err error) {
	return nil
}
