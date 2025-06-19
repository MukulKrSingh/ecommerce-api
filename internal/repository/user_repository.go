package repository

import (
	"github.com/ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) (model.User, error)
	Login(email string, password string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id string) (model.User, error)
	UpdateUser(user *model.User) (model.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) (model.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}
func (r *userRepository) Login(email string, password string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(id string) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *model.User) (model.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return model.User{}, err
	}
	return *user, nil
}

func (r *userRepository) DeleteUser(id string) error {

	return r.db.Delete(&model.User{}, id).Error
}
