package service

import (
	"go-on-store/app/model"
	"go-on-store/app/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) GetTransactionByID(id uint) (*model.Transaction, error) {
	return s.repo.GetTransactionByID(id)
}

func (s *TransactionService) AddTransaction(transaction *model.Transaction) error {
	return s.repo.AddTransaction(transaction)
}
