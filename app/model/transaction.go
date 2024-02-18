package model

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Total  float64 `json:"total"`
}
