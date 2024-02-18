package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index" json:"username"`
	Password string `gorm:"size:100" json:"password"`
}
