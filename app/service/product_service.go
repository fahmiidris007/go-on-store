package service

import (
	"go-on-store/app/model"
	"go-on-store/app/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) GetProductByCategoryID(categoryID uint) ([]*model.Product, error) {
	return s.repo.GetProductByCategoryID(categoryID)
}

func (s *ProductService) AddProduct(product *model.Product) error {
	return s.repo.AddProduct(product)
}
