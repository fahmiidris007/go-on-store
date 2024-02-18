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

func (s *ProductService) GetAllProducts() ([]*model.Product, error) {
	return s.repo.GetAllProducts()
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

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}
