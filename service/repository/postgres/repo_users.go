package postgres

import (
	"fmt"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usersRepository {
	return usersRepository{
		db: db,
	}
}

func (r usersRepository) CreateUser(user models.Users) (userId uint, err error) {
	err = r.db.Create(&user).Error
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return user.ID, err
	}
	return user.ID, nil
}

func (r usersRepository) GetUserById(userId uint) (user models.Users, err error) {
	err = r.db.Debug().Where("id = ?", userId).Take(&user).Error
	if err != nil {
		fmt.Printf("[UserRepository][GetUserByEmail] error while check user by email: %+v", err)
		return user, err
	}
	return user, err
}

func (r usersRepository) GetUserByEmail(email string) (user models.Users, err error) {
	err = r.db.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		fmt.Printf("[UserRepository][GetUserByEmail] error while check user by email: %+v", err)
		return user, err
	}
	return user, err
}

func (r usersRepository) EmailIsExist(email string) (db *gorm.DB) {
	result := r.db.Where("email = ?", email).Find(&models.Users{})

	return result
}
