package service

import (
	"github.com/ecommerce-api/internal/model"
	"github.com/ecommerce-api/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(user *model.User) (model.User, error) {
	return s.userRepo.CreateUser(user)
}
func (s *UserService) Login(email string, password string) (model.User, error) {
	return s.userRepo.Login(email, password)
}

func (s *UserService) GetUserByEmail(email string) (model.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *model.User) (model.User, error) {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}
