package service

import (
	"github.com/Fachrulmustofa20/bank-example.git/models"
)

type UsersUsecase interface {
	Register(registerReq models.RegisterRequest) error
	Login(email string, password string) (token string, err error)
	Profile(id uint) (response models.Users, err error)
}

type BalanceUsecase interface {
	GetBalance(userId uint) (balance models.Balance, err error)
	TopUpBalance(topUp models.TopUpRequest, userId uint) (err error)
	TransferBalance(transfer models.TransferBalance, userId uint) (err error)
	GetMutationBalance(userId uint) (history []models.BalanceHistory, err error)
}

type BankUsecase interface {
	CreateAccountBank(bank models.Bank) (err error)
	GetBalanceBankByCode(code string) (balanceInBank models.Bank, err error)
	AddDeposit(bank models.Bank) (err error)
}
