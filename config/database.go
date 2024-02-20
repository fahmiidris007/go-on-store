package config

import (
	"go-on-store/app/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPassword + "@tcp(db-mysql:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{}, &model.Cart{}, &model.Transaction{})

	// if db not exists, seed the database
	if db.First(&model.Product{}).RecordNotFound() {
		Seed(db)
	}

	return db
}

func Seed(db *gorm.DB) {
	user := model.User{
		Username: "user1",
		Password: "password1",
	}
	db.Create(&user)

	product := model.Product{
		Name:       "product1",
		Price:      1000,
		CategoryID: 1,
	}
	db.Create(&product)

	category := model.Category{
		Name: "category1",
	}
	db.Create(&category)

	cart := model.Cart{
		UserID:    user.ID,
		ProductID: product.ID,
		Quantity:  1,
	}
	db.Create(&cart)

	transaction := model.Transaction{
		UserID: user.ID,
		Total:  1000,
	}
	db.Create(&transaction)

}
