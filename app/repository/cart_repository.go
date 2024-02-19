package repository

import (
	"go-on-store/app/model"

	"github.com/jinzhu/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) GetCartByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	if err := r.DB.First(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) AddCart(cart *model.Cart) error {
	return r.DB.Create(cart).Error
}

func (r *CartRepository) UpdateCart(cart *model.Cart) error {
	return r.DB.Save(cart).Error
}

func (r *CartRepository) GetCartByUserID(userID uint) ([]*model.Cart, error) {
	var carts []*model.Cart
	if err := r.DB.Where("user_id = ?", userID).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepository) DeleteCart(cart *model.Cart) error {
	return r.DB.Delete(cart).Error
}
