package usecase

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Fachrulmustofa20/bank-example.git/constants"
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
)

type bankUsecase struct {
	bankRepo service.BankRepository
	userRepo service.UsersRepository
}

func NewBankUsecase(bankRepo service.BankRepository, userRepo service.UsersRepository) service.BankUsecase {
	return &bankUsecase{
		bankRepo: bankRepo,
		userRepo: userRepo,
	}
}

func (usecase bankUsecase) CreateAccountBank(bank models.Bank) (err error) {
	bankId, err := usecase.bankRepo.CreateAccountBank(bank)
	if err != nil {
		log.Error("error create account bank: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	user, err := usecase.userRepo.GetUserById(bank.UserId)
	if err != nil {
		log.Error("error get user by id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// auto create history bank
	ip := utils.GetLocalIP()
	historyBank := models.BankBalanceHistory{
		BankBalanceId: bankId,
		BalanceBefore: 0,
		BalanceAfter:  bank.Balance,
		Activity:      "Create Account Bank",
		Type:          "debit",
		Ip:            ip,
		Location:      "Jawa Tengah",
		UserAgent:     "Postman 3.0",
		Author:        user.Username,
	}
	err = usecase.bankRepo.CreateHistoryInBank(historyBank)
	if err != nil {
		log.Error("error create history bank: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase bankUsecase) GetBalanceBankByCode(code string) (balanceInBank models.Bank, err error) {
	balanceInBank, err = usecase.bankRepo.GetBalanceBankByCode(code)
	if err != nil {
		log.Error("error get balance bank by code: ", err)
		err = fmt.Errorf("%+v", err)
		return balanceInBank, err
	}
	return balanceInBank, err
}

func (usecase bankUsecase) AddDeposit(bank models.Bank) (err error) {
	currentBalance, err := usecase.bankRepo.GetBalanceBankByCode(bank.Code)
	if err != nil {
		log.Error("error get balance bank by code: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	if currentBalance.Code != bank.Code {
		log.Info("code account bank not same")
		return constants.ErrCodeBank
	}

	// update balance in bank
	balanceTotal := currentBalance.Balance + bank.Balance
	bankUpdate := models.Bank{
		Balance:        balanceTotal,
		BalanceAchieve: balanceTotal,
		UserId:         bank.UserId,
	}
	err = usecase.bankRepo.UpdateBalanceBankByUserId(bankUpdate)
	if err != nil {
		log.Error("error update balance bank by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// get Ip
	ip := utils.GetLocalIP()

	// get user
	user, err := usecase.userRepo.GetUserById(bank.UserId)
	if err != nil {
		log.Error("error get user by id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// add history bank
	bankHistory := models.BankBalanceHistory{
		BankBalanceId: currentBalance.ID,
		BalanceBefore: balanceTotal,
		BalanceAfter:  balanceTotal,
		Activity:      "Add Deposit",
		Type:          "debit",
		Ip:            ip,
		Location:      "Jawa Tengah",
		UserAgent:     "Postman",
		Author:        user.Username,
	}
	err = usecase.bankRepo.CreateHistoryInBank(bankHistory)
	if err != nil {
		log.Error("error create history bank: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}
