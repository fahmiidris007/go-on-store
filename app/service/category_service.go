package service

import (
	"go-on-store/app/model"
	"go-on-store/app/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]*model.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repo.GetCategoryByID(id)
}

func (s *CategoryService) AddCategory(category *model.Category) error {
	return s.repo.AddCategory(category)
}

func (s *CategoryService) UpdateCategory(category *model.Category) error {
	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}
