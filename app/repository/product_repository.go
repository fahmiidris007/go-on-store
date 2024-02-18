package repository

import (
	"go-on-store/app/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) AddProduct(product *model.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) GetProductByCategoryID(categoryID uint) ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
