package postgres

import (
	"fmt"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) bankRepository {
	return bankRepository{db: db}
}

func (r bankRepository) CreateAccountBank(bank models.Bank) (bankId uint, err error) {
	err = r.db.Create(&bank).Error
	if err != nil {
		log.Error("error while create account bank: ", err)
		err = fmt.Errorf("%+v", err)
		return bank.ID, err
	}
	return bank.ID, nil
}

func (r bankRepository) GetBalanceBankByCode(code string) (balanceInBank models.Bank, err error) {
	err = r.db.Debug().Where("code = ?", code).Take(&balanceInBank).Error
	if err != nil {
		log.Error("error while balance bank by code: ", err)
		return balanceInBank, err
	}
	return balanceInBank, nil
}

func (r bankRepository) GetBalanceBankByUserId(userId uint) (balanceInBank models.Bank, err error) {
	err = r.db.Debug().Where("user_id = ?", userId).Take(&balanceInBank).Error
	if err != nil {
		log.Error("error while balance bank by userId: ", err)
		return balanceInBank, err
	}
	return balanceInBank, nil
}

func (r bankRepository) UpdateBalanceByCode(bank models.Bank) (err error) {
	err = r.db.Model(&bank).Where("code = ?", bank.Code).Update("balance", bank.Balance).Error
	if err != nil {
		log.Error("error while update bank balance by code: ", err)
		return err
	}
	return err
}

func (r bankRepository) CreateHistoryInBank(bankHistory models.BankBalanceHistory) (err error) {
	err = r.db.Create(&bankHistory).Error
	if err != nil {
		log.Error("error while create history in bank: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}
	return nil
}

func (r bankRepository) UpdateBalanceBankByUserId(bank models.Bank) (err error) {
	err = r.db.Debug().Model(bank).Where("user_id = ?", bank.UserId).Updates(models.Bank{
		Balance:        bank.Balance,
		BalanceAchieve: bank.BalanceAchieve,
	}).Error

	if err != nil {
		log.Error("error while update balance bank by userId: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}
	return nil
}
