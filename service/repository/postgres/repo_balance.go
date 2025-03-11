package postgres

import (
	"fmt"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) balanceRepository {
	return balanceRepository{
		db: db,
	}
}

func (r balanceRepository) GetBalance(userID uint) (balance models.Balance, err error) {
	err = r.db.Debug().Where("user_id = ?", userID).Take(&balance).Error
	if err != nil {
		log.Error("error while get balance by userId: ", err)
		return balance, err
	}
	return balance, nil
}

func (r balanceRepository) UpdateUserBalanceByUserId(userBalance models.Balance, userId uint) (err error) {
	err = r.db.Debug().Model(userBalance).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"balance":         userBalance.Balance,
		"balance_achieve": userBalance.BalanceAchieve,
	}).Error
	if err != nil {
		log.Error("error while update user balance by userId: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (r balanceRepository) CreateUserBalance(userBalance models.Balance) (err error) {
	err = r.db.Create(&userBalance).Error
	if err != nil {
		log.Error("error while create user balance: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}
	return nil
}

func (r balanceRepository) CreateBalanceHistory(balanceHistory models.BalanceHistory) (err error) {
	err = r.db.Create(&balanceHistory).Error
	if err != nil {
		log.Error("error while create balance history: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}
	return nil
}

func (r balanceRepository) GetBalanceHistoryByUser(author string) (balanceHistory []models.BalanceHistory, err error) {
	err = r.db.Where("author = ?", author).Find(&balanceHistory).Error
	if err != nil {
		log.Error("error while get balance History by author: ", err)
		return balanceHistory, err
	}
	return balanceHistory, nil
}

func (r balanceRepository) GetBalanceHistoryByBalanceID(userBalanceId uint) (balanceHistory []models.BalanceHistory, err error) {
	err = r.db.Debug().Where("user_balance_id = ?", userBalanceId).Order("id DESC").Find(&balanceHistory).Error
	if err != nil {
		log.Error("error while get balance History by author: ", err)
		return balanceHistory, err
	}
	return balanceHistory, nil
}
