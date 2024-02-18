package service

import (
	"go-on-store/app/model"
	"go-on-store/app/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *UserService) AddUser(user *model.User) error {
	return s.repo.AddUser(user)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}
