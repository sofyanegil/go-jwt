package services

import (
	"errors"
	"fmt"
	"go-jwt/models"
	"go-jwt/repository"

	"github.com/stretchr/testify/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{ProductRepository: productRepository}

func TestReadFound(t *testing.T) {
	product := models.Product{
		Title:       "Title testify",
		Description: "Description testify",
		UserID:      1,
	}

	productRepository.Mock.On("ReadProduct", 1).Return(&product, nil)

	result, err := productService.ReadProduct(1)

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Title, result.Title, "result has to be 'Title testify'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Description testify'")
	assert.Equal(t, &product, result, "result has to be a product data with id '1'")
}

func TestReadNotFound(t *testing.T) {
	productRepository.Mock.On("ReadProduct", 2).Return(nil, errors.New("data doesn't exist"))

	result, err := productService.ReadProduct(2)

	assert.Nil(t, result)

	assert.NotNil(t, err)
	assert.Equal(t, "data doesn't exist", err.Error(), "error response has to be 'data doesn't exist'")
}

func TestReadAllFound(t *testing.T) {
	products := &[]models.Product{
		{Title: "Title testify0", Description: "Description testify0", UserID: 1},
		{Title: "Title testify1", Description: "Description testify1", UserID: 1},
	}

	productRepository.Mock.On("ReadAllProduct").Return(products, nil).Once()

	results, err := productService.ReadAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, &results)

	assert.Equal(t, products, results, "result has to be a products data with id '1'")

	for i, result := range *results {
		assert.Equal(t, (*products)[i].Title, result.Title, fmt.Sprintf("result has to be 'Title testify%d'", i))
		assert.Equal(t, (*products)[i].Description, result.Description, fmt.Sprintf("result has to be 'Description testify%d'", i))
	}
}

func TestReadAllNotFound(t *testing.T) {
	productRepository.Mock.On("ReadAllProduct").Return(nil, errors.New("data doesn't exist")).Once()

	result, err := productService.ReadAllProduct()

	assert.Nil(t, result)

	assert.NotNil(t, err)
	assert.Equal(t, "data doesn't exist", err.Error(), "error response has to be 'data doesn't exist'")
}
