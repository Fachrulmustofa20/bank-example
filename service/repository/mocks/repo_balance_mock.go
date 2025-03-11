package mocks

import (
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/stretchr/testify/mock"
)

type BalanceRepository struct {
	mock.Mock
}

func (m *BalanceRepository) GetBalance(userID uint) (balance models.Balance, err error) {
	args := m.Called(userID)
	return args.Get(0).(models.Balance), args.Error(1)
}

func (m *BalanceRepository) UpdateUserBalanceByUserId(userBalance models.Balance, userId uint) (err error) {
	args := m.Called(userBalance, userId)
	return args.Error(0)
}

func (m *BalanceRepository) CreateUserBalance(userBalance models.Balance) (err error) {
	args := m.Called(userBalance)
	return args.Error(0)
}

func (m *BalanceRepository) CreateBalanceHistory(balanceHistory models.BalanceHistory) (err error) {
	args := m.Called(balanceHistory)
	return args.Error(0)
}

func (m *BalanceRepository) GetBalanceHistoryByUser(author string) (balanceHistory []models.BalanceHistory, err error) {
	args := m.Called(author)
	return args.Get(0).([]models.BalanceHistory), args.Error(1)
}

func (m *BalanceRepository) GetBalanceHistoryByBalanceID(UserBalanceId uint) (balanceHistory []models.BalanceHistory, err error) {
	args := m.Called(UserBalanceId)
	return args.Get(0).([]models.BalanceHistory), args.Error(1)
}
