package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
}
