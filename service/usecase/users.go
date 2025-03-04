package usecase

import (
	"fmt"

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
		err = fmt.Errorf("%+v", err)
		return err
	}

	return nil
}

func (usecase usersUsecase) Login(email string, password string) (token string, err error) {
	users, err := usecase.userRepo.GetUserByEmail(email)
	if err != nil {
		return token, constants.ErrLogin
	}

	comparePass := utils.ComparePassword([]byte(users.Password), []byte(password))
	if !comparePass {
		return token, constants.ErrLogin
	}

	token, err = utils.GenerateToken(users.ID, users.Email)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (usecase usersUsecase) Profile(id uint) (response models.Users, err error) {
	users, err := usecase.userRepo.GetUserById(id)
	users.Password = ""
	if err != nil {
		return response, err
	}
	return users, err
}
