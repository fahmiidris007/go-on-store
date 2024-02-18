package service

import (
	"go-on-store/app/model"
	"go-on-store/app/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) GetCartByID(id uint) (*model.Cart, error) {
	return s.repo.GetCartByID(id)
}

func (s *CartService) GetCartByUserID(userID uint) ([]*model.Cart, error) {
	return s.repo.GetCartByUserID(userID)
}

func (s *CartService) AddCart(cart *model.Cart) error {
	return s.repo.AddCart(cart)
}

func (s *CartService) DeleteCart(cart *model.Cart) error {
	return s.repo.DeleteCart(cart)
}
