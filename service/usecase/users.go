package usecase

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Fachrulmustofa20/bank-example.git/constants"
	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
)

type usersUsecase struct {
	userRepo    service.UsersRepository
	balanceRepo service.BalanceRepository
}

func NewUsersUsecase(userRepo service.UsersRepository, balanceRepo service.BalanceRepository) service.UsersUsecase {
	return &usersUsecase{
		userRepo:    userRepo,
		balanceRepo: balanceRepo,
	}
}

func (usecase usersUsecase) Register(registerReq models.RegisterRequest) error {
	// hash password
	password := utils.HashPass(registerReq.Password)
	userId, err := usecase.userRepo.CreateUser(models.Users{
		Username: registerReq.Username,
		Password: password,
		Email:    registerReq.Email,
	})
	if err != nil {
		log.Error("error create user: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	// auto create user balance
	userBalance := models.Balance{
		UserId:         userId,
		Balance:        0,
		BalanceAchieve: 0,
	}
	err = usecase.balanceRepo.CreateUserBalance(userBalance)
	if err != nil {
		log.Error("error create user balance: ", err)
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase usersUsecase) Login(email string, password string) (token string, err error) {
	users, err := usecase.userRepo.GetUserByEmail(email)
	if err != nil {
		log.Error("error get user by email: ", err)
		return token, constants.ErrLogin
	}

	comparePass := utils.ComparePassword([]byte(users.Password), []byte(password))
	if !comparePass {
		log.Info("password not match")
		return token, constants.ErrLogin
	}

	token, err = utils.GenerateToken(users.ID, users.Email)
	if err != nil {
		log.Error("error generate token: ", err)
		return token, err
	}

	return token, nil
}

func (usecase usersUsecase) Profile(id uint) (response models.Users, err error) {
	users, err := usecase.userRepo.GetUserById(id)
	users.Password = ""
	if err != nil {
		log.Error("error get user by id: ", err)
		return response, err
	}
	return users, err
}
