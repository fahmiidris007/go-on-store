package repository

import (
	"go-on-store/app/model"

	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAllProducts() ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) AddProduct(product *model.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetProductByCategoryID(categoryID uint) ([]*model.Product, error) {
	var products []*model.Product
	if err := r.DB.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&model.Product{}, id).Error
}
