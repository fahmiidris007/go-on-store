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

func (r *TransactionRepository) GetTransactionByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) AddTransaction(transaction *model.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r *TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	return r.DB.Create(transaction).Error
}
