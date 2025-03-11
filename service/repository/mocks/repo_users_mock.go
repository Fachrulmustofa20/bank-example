package mocks

import (
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UsersRepository struct {
	mock.Mock
}

func (m *UsersRepository) CreateUser(user models.Users) (uint, error) {
	args := m.Called(user)
	return args.Get(0).(uint), args.Error(1)
}

func (m *UsersRepository) GetUserByEmail(email string) (models.Users, error) {
	args := m.Called(email)
	return args.Get(0).(models.Users), args.Error(1)
}

func (m *UsersRepository) GetUserById(id uint) (models.Users, error) {
	args := m.Called(id)
	return args.Get(0).(models.Users), args.Error(1)
}

func (m *UsersRepository) EmailIsExist(email string) (db *gorm.DB) {
	args := m.Called(email)
	return args.Get(0).(*gorm.DB)
}
