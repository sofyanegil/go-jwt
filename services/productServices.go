package services

import (
	"go-jwt/models"
	"go-jwt/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepository repository.IProductRepository
	DB                *gorm.DB
}

func NewProductService(productRepository repository.IProductRepository, db *gorm.DB) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
		DB:                db,
	}
}

func (service ProductService) CreateProduct(p *models.Product, userID uint) (product *models.Product, err error) {
	p.UserID = userID

	product, err = service.ProductRepository.CreateProduct(service.DB, p)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (service ProductService) ReadAllProduct() (product *[]models.Product, err error) {
	product, err = service.ProductRepository.ReadAllProduct(service.DB)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service ProductService) ReadProduct(productID int) (product *models.Product, err error) {
	product, err = service.ProductRepository.ReadProduct(service.DB, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service ProductService) UpdateProduct(p *models.Product, productID int) (product *models.Product, err error) {
	product, err = service.ProductRepository.UpdateProduct(service.DB, p, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service ProductService) DeleteProduct(productID int) (err error) {
	err = service.ProductRepository.DeleteProduct(service.DB, productID)
	if err != nil {
		return err
	}
	return nil
}
