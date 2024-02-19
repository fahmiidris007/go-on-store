package repository

import (
	"go-on-store/app/model"

	"github.com/jinzhu/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetTransactionByUserID(userID uint) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) AddTransaction(transaction *model.Transaction) error {
	return r.DB.Create(transaction).Error
}
