package repository

import (
	"errors"
	"go-jwt/models"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (repository ProductRepository) CreateProduct(db *gorm.DB, p *models.Product) (product *models.Product, err error) {
	err = db.Debug().Create(&p).Error
	product = p
	return product, err
}

func (repository ProductRepository) ReadAllProduct(db *gorm.DB) (product *[]models.Product, err error) {
	err = db.Debug().Find(&product).Error
	return product, err
}

func (repository ProductRepository) ReadProduct(db *gorm.DB, productID int) (product *models.Product, err error) {
	err = db.Debug().First(&product, "id = ?", productID).Error
	return product, err
}

func (repository ProductRepository) UpdateProduct(db *gorm.DB, p *models.Product, productID int) (product *models.Product, err error) {
	dbreturn := db.Model(&p).Clauses(clause.Returning{}).Where("id = ?", productID).Updates(models.Product{Title: p.Title, Description: p.Description})
	if dbreturn.Error != nil {
		return nil, dbreturn.Error
	}

	if dbreturn.RowsAffected == 0 {
		return nil, errors.New("data doesn't exist")
	}
	return p, nil
}

func (repository ProductRepository) DeleteProduct(db *gorm.DB, productID int) (err error) {
	dbreturn := db.Where("id = ?", productID).Delete(&models.Product{})
	if dbreturn.Error != nil {
		return dbreturn.Error
	}

	if dbreturn.RowsAffected == 0 {
		return errors.New("data doesn't exist")
	}
	return nil
}
