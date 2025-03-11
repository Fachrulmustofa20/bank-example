package mocks

import (
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/stretchr/testify/mock"
)

type BankRepository struct {
	mock.Mock
}

func (m *BalanceRepository) CreateAccountBank(bank models.Bank) (bankId uint, err error) {
	args := m.Called(bank)
	return args.Get(0).(uint), args.Error(1)
}

func (m *BalanceRepository) GetBalanceBankByCode(code string) (balanceInBank models.Bank, err error) {
	args := m.Called(code)
	return args.Get(0).(models.Bank), args.Error(1)
}

func (m *BalanceRepository) GetBalanceBankByUserId(userId uint) (balanceInBank models.Bank, err error) {
	args := m.Called(userId)
	return args.Get(0).(models.Bank), args.Error(1)
}

func (m *BalanceRepository) UpdateBalanceByCode(bank models.Bank) (err error) {
	args := m.Called(bank)
	return args.Error(0)
}

func (m *BalanceRepository) CreateHistoryInBank(bankHistory models.BankBalanceHistory) (err error) {
	args := m.Called(bankHistory)
	return args.Error(0)
}

func (m *BalanceRepository) UpdateBalanceBankByUserId(bank models.Bank) (err error) {
	args := m.Called(bank)
	return args.Error(0)
}
