package repository

import (
	"go-on-store/app/model"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) AddCategory(category *model.Category) error {
	return r.DB.Save(category).Error
}
