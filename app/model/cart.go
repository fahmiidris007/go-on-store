package model

import (
	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
