package repository

import (
	"go-on-store/app/model"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) GetAllCategories() ([]*model.Category, error) {
	var categories []*model.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil

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

func (r *CategoryRepository) UpdateCategory(category *model.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepository) DeleteCategory(id uint) error {
	return r.DB.Delete(&model.Category{}, id).Error
}
