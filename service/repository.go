package service

import (
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateUser(user models.Users) (userId uint, err error)
	GetUserByEmail(email string) (user models.Users, err error)
	GetUserById(userId uint) (user models.Users, err error)
	EmailIsExist(email string) (db *gorm.DB)
}

type BalanceRepository interface {
	GetBalance(userID uint) (balance models.Balance, err error)
	UpdateUserBalanceByUserId(userBalance models.Balance, userId uint) (err error)
	CreateUserBalance(userBalance models.Balance) (err error)
	CreateBalanceHistory(balanceHistory models.BalanceHistory) (err error)
	GetBalanceHistoryByUser(author string) (balanceHistory []models.BalanceHistory, err error)
	GetBalanceHistoryByBalanceID(UserBalanceId uint) (balanceHistory []models.BalanceHistory, err error)
}

type BankRepository interface {
	CreateAccountBank(bank models.Bank) (bankId uint, err error)
	GetBalanceBankByCode(code string) (balanceInBank models.Bank, err error)
	GetBalanceBankByUserId(userId uint) (balanceInBank models.Bank, err error)
	UpdateBalanceByCode(bank models.Bank) (err error)
	CreateHistoryInBank(bankHistory models.BankBalanceHistory) (err error)
	UpdateBalanceBankByUserId(bank models.Bank) (err error)
}
