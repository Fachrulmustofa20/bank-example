package usecase

import (
	"errors"
	"fmt"

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
	balanceInBank, err := usecase.bankRepo.GetBalanceBankByCode(topUp.CodeAccountBank)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return err
	}

	// validation userid not same
	if userId != balanceInBank.UserId {
		return constants.ErrCodeBank
	}

	// validation balance in bank
	if balanceInBank.Balance < topUp.Amount {
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
		err = fmt.Errorf("%+v", err)
		return err
	}

	// get Ip
	ip := utils.GetLocalIP()

	// get user
	user, err := usecase.userRepo.GetUserById(userId)
	if err != nil {
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
		err = fmt.Errorf("%+v", err)
		return err
	}

	// update user balance
	balance, err := usecase.balanceRepo.GetBalance(userId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return err
	}
	totalUserBalance := balance.Balance + topUp.Amount
	fmt.Print(totalUserBalance)
	balanceUpdate := models.Balance{
		Balance:        totalUserBalance,
		BalanceAchieve: totalUserBalance,
		UserId:         userId,
	}
	err = usecase.balanceRepo.UpdateUserBalanceByUserId(balanceUpdate, userId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return err
	}

	// create transaction history in user balance
	historyBalance := models.BalanceHistory{
		UserBalanceId: balance.ID,
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
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase balanceUsecase) TransferBalance(transfer models.TransferBalance, userId uint) (err error) {
	// validate email user receipent
	tx := usecase.userRepo.EmailIsExist(transfer.EmailRecipient)
	if tx.RowsAffected < 1 {
		return errors.New("recipient emails not found")
	}

	user, err := usecase.userRepo.GetUserById(userId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return err
	}

	balanceSender, err := usecase.balanceRepo.GetBalance(userId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return err
	}

	if balanceSender.Balance < transfer.Amount {
		return errors.New("your balance is insufficient. please top up first")
	}

	// validate email
	if user.Email == transfer.EmailRecipient {
		return errors.New("recipient emails cannot be the same")
	}

	// deduct user balance sender
	totalAmountBalanceSender := balanceSender.Balance - transfer.Amount
	deductBalanceSender := models.Balance{
		Balance:        totalAmountBalanceSender,
		BalanceAchieve: balanceSender.Balance,
	}
	err = usecase.balanceRepo.UpdateUserBalanceByUserId(deductBalanceSender, userId)
	if err != nil {
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
		err = fmt.Errorf("%+v", err)
		return err
	}

	// update amount user recepient
	userRecepient, err := usecase.userRepo.GetUserByEmail(transfer.EmailRecipient)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return err
	}
	balanceRecepient, err := usecase.balanceRepo.GetBalance(userRecepient.ID)
	if err != nil {
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
		err = fmt.Errorf("%+v", err)
		return err
	}

	// create transaction history in user balance
	historyBalance := models.BalanceHistory{
		UserBalanceId: balanceRecepient.ID,
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
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase balanceUsecase) GetMutationBalance(userId uint) (history []models.BalanceHistory, err error) {
	balance, err := usecase.GetBalance(userId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return nil, err
	}

	balanceHistory, err := usecase.balanceRepo.GetBalanceHistoryByBalanceID(balance.UserId)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return nil, err
	}

	return balanceHistory, nil
}
