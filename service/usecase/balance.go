package usecase

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Fachrulmustofa20/bank-example.git/constants"
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
)

type balanceUsecase struct {
	userRepo    service.UsersRepository
	balanceRepo service.BalanceRepository
	bankRepo    service.BankRepository
}

func NewBalanceUsecase(userRepo service.UsersRepository, balanceRepo service.BalanceRepository, bankRepo service.BankRepository) service.BalanceUsecase {
	return &balanceUsecase{
		userRepo:    userRepo,
		balanceRepo: balanceRepo,
		bankRepo:    bankRepo,
	}
}

func (usecase balanceUsecase) GetBalance(userId uint) (balance models.Balance, err error) {
	balance, err = usecase.balanceRepo.GetBalance(userId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return balance, err
	}
	return balance, nil
}

func (usecase balanceUsecase) TopUpBalance(topUp models.TopUpRequest, userId uint) (err error) {
	balanceInBank, err := usecase.bankRepo.GetBalanceBankByUserId(userId)
	if err != nil {
		log.Error("error get balance bank by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// validation userid not same
	if topUp.CodeAccountBank != balanceInBank.Code {
		log.Info("code account bank not same")
		return constants.ErrCodeBank
	}

	// validation balance in bank
	if balanceInBank.Balance < topUp.Amount {
		log.Info("balance in bank is not enough")
		return constants.ErrBalanceNotEnough
	}

	// update bank balance
	totalBalanceInBank := balanceInBank.Balance - topUp.Amount
	updateBalance := models.Bank{
		Balance: totalBalanceInBank,
		Code:    topUp.CodeAccountBank,
	}
	err = usecase.bankRepo.UpdateBalanceByCode(updateBalance)
	if err != nil {
		log.Error("error update balance by code: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// get Ip
	ip := utils.GetLocalIP()

	// get user
	user, err := usecase.userRepo.GetUserById(userId)
	if err != nil {
		log.Error("error get user by id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// create history in bank
	historyBank := models.BankBalanceHistory{
		BankBalanceId: balanceInBank.ID,
		BalanceBefore: balanceInBank.Balance,
		BalanceAfter:  totalBalanceInBank,
		Activity:      "Top Up Balance User",
		Type:          "kredit",
		Ip:            ip,
		Location:      "Jawa Tengah",
		UserAgent:     "Postman 3.0",
		Author:        user.Username,
	}
	err = usecase.bankRepo.CreateHistoryInBank(historyBank)
	if err != nil {
		log.Error("error create history in bank: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// update user balance
	balance, err := usecase.balanceRepo.GetBalance(userId)
	if err != nil {
		log.Error("error get balance by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}
	totalUserBalance := balance.Balance + topUp.Amount
	balanceUpdate := models.Balance{
		Balance:        totalUserBalance,
		BalanceAchieve: totalUserBalance,
		UserId:         userId,
	}
	err = usecase.balanceRepo.UpdateUserBalanceByUserId(balanceUpdate, userId)
	if err != nil {
		log.Error("error update balance by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// create transaction history in user balance
	historyBalance := models.BalanceHistory{
		UserBalanceId: userId,
		BalanceBefore: balance.Balance,
		BalanceAfter:  totalUserBalance,
		Activity:      "Top Up Balance",
		Type:          "debit",
		Ip:            ip,
		Location:      "Jawa Tengah",
		UserAgent:     "Postman 3.0",
		Author:        user.Username,
	}
	err = usecase.balanceRepo.CreateBalanceHistory(historyBalance)
	if err != nil {
		log.Error("error create balance history: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase balanceUsecase) TransferBalance(transfer models.TransferBalance, userId uint) (err error) {
	// validate email user receipent
	tx := usecase.userRepo.EmailIsExist(transfer.EmailRecipient)
	if tx.RowsAffected < 1 {
		log.Info("recipient emails not found")
		return errors.New("recipient emails not found")
	}

	user, err := usecase.userRepo.GetUserById(userId)
	if err != nil {
		log.Error("error get user by id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	balanceSender, err := usecase.balanceRepo.GetBalance(userId)
	if err != nil {
		log.Error("error get balance by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	if balanceSender.Balance < transfer.Amount {
		log.Info("balance is insufficient")
		return errors.New("your balance is insufficient. please top up first")
	}

	// validate email
	if user.Email == transfer.EmailRecipient {
		log.Info("recipient emails cannot be the same")
		return errors.New("recipient emails cannot be the same")
	}

	// deduct user balance sender
	totalAmountBalanceSender := balanceSender.Balance - transfer.Amount
	deductBalanceRecepient := models.Balance{
		Balance:        totalAmountBalanceSender,
		BalanceAchieve: totalAmountBalanceSender,
	}
	err = usecase.balanceRepo.UpdateUserBalanceByUserId(deductBalanceRecepient, user.ID)
	if err != nil {
		log.Error("error update balance by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// get Ip
	ip := utils.GetLocalIP()

	// create history user balance sender
	historyBalanceSender := models.BalanceHistory{
		UserBalanceId: userId,
		BalanceBefore: balanceSender.Balance,
		BalanceAfter:  totalAmountBalanceSender,
		Activity:      "Transfer Balance",
		Type:          "kredit",
		Ip:            ip,
		Location:      "Jawa Tengah",
		UserAgent:     "Postman 3.0",
		Author:        user.Username,
	}
	err = usecase.balanceRepo.CreateBalanceHistory(historyBalanceSender)
	if err != nil {
		log.Error("error create balance history: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// update amount user recepient
	userRecepient, err := usecase.userRepo.GetUserByEmail(transfer.EmailRecipient)
	if err != nil {
		log.Error("error get user by email: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}
	balanceRecepient, err := usecase.balanceRepo.GetBalance(userRecepient.ID)
	if err != nil {
		log.Error("error get balance by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	totalAmountBalanceRecepient := balanceRecepient.Balance + transfer.Amount
	addBalanceRecepient := models.Balance{
		Balance:        totalAmountBalanceRecepient,
		BalanceAchieve: totalAmountBalanceRecepient,
	}
	err = usecase.balanceRepo.UpdateUserBalanceByUserId(addBalanceRecepient, userRecepient.ID)
	if err != nil {
		log.Error("error update balance by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// create transaction history in user balance
	historyBalance := models.BalanceHistory{
		UserBalanceId: userRecepient.ID,
		BalanceBefore: balanceRecepient.Balance,
		BalanceAfter:  totalAmountBalanceRecepient,
		Activity:      "Transfer Balance",
		Type:          "debit",
		Ip:            ip,
		Location:      "Jawa Tengah",
		UserAgent:     "Postman 3.0",
		Author:        user.Username,
	}
	err = usecase.balanceRepo.CreateBalanceHistory(historyBalance)
	if err != nil {
		log.Error("error create balance history: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase balanceUsecase) GetMutationBalance(userId uint) (history []models.BalanceHistory, err error) {
	balanceHistory, err := usecase.balanceRepo.GetBalanceHistoryByBalanceID(userId)
	if err != nil {
		log.Error("error get balance history by user id: ", err)
		err = fmt.Errorf("%+v", err)
		return nil, err
	}

	return balanceHistory, nil
}
