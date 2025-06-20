package service

import (
	"errors"

	"github.com/ecommerce-api/internal/helpers"
	"github.com/ecommerce-api/internal/model"
	"github.com/ecommerce-api/internal/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
	Auth     helpers.Auth
}

func NewUserService(UserRepo repository.UserRepository, auth helpers.Auth) *UserService {
	return &UserService{UserRepo: UserRepo, Auth: auth}
}

func (s *UserService) RegisterUser(user *model.User) (string, model.User, error) {

	//Hash the password
	hashedPassword, err := s.Auth.CreateHashedPassword(user.Password)
	if err != nil {
		return "", model.User{}, err
	}
	user.Password = hashedPassword

	//Create the user
	u, err := s.UserRepo.CreateUser(&model.User{
		// Auto-incremented by the database
		Email:    user.Email,
		Password: hashedPassword,
		Role:     user.Role,
	})

	if err != nil {
		return "", model.User{}, err
	}

	//Generate JWT token
	token, err := s.Auth.GenerateToken(u.ID, u.Email, u.Role)
	if err != nil {
		return "", model.User{}, err
	}

	return token, *user, nil
}
func (s *UserService) Login(email string, password string) (string, error) {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	//Check if user exists
	if user.ID == 0 {
		return "", errors.New("user not found")
	}

	//Verify password
	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	//Generate JWT token
	return s.Auth.GenerateToken(user.ID, user.Email, user.Role)

}

func (s *UserService) GetUserByEmail(email string) (model.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *model.User) (model.User, error) {
	return s.UserRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.UserRepo.DeleteUser(id)
}
