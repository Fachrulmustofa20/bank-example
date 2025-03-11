package usecase_test

import (
	"testing"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service/repository/mocks"
	"github.com/Fachrulmustofa20/bank-example.git/service/usecase"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterSuccess(t *testing.T) {
	mockUserRepo := new(mocks.UsersRepository)
	mockBalanceRepo := new(mocks.BalanceRepository)
	usecase := usecase.NewUsersUsecase(mockUserRepo, mockBalanceRepo)

	request := models.RegisterRequest{
		Username: "test",
		Password: "test",
		Email:    "test@gmail.com",
	}

	mockUserRepo.On("CreateUser", mock.Anything).Return(uint(1), nil)
	mockBalanceRepo.On("CreateUserBalance", mock.Anything).Return(nil)

	err := usecase.Register(request)
	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}

func TestLoginSuccess(t *testing.T) {
	mockUserRepo := new(mocks.UsersRepository)
	mockBalanceRepo := new(mocks.BalanceRepository)
	usecase := usecase.NewUsersUsecase(mockUserRepo, mockBalanceRepo)

	reqPassword := "test"
	reqEmail := "test"
	hashPassword := utils.HashPass(reqPassword)

	mockUserRepo.On("GetUserByEmail", mock.Anything).Return(models.Users{
		Username: "test",
		Password: hashPassword,
		Email:    "test@gmail.com",
	}, nil)

	token, err := usecase.Login(reqEmail, reqPassword)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockUserRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}

func TestProfileSuccess(t *testing.T) {
	mockUserRepo := new(mocks.UsersRepository)
	mockBalanceRepo := new(mocks.BalanceRepository)
	usecase := usecase.NewUsersUsecase(mockUserRepo, mockBalanceRepo)

	reqId := 1
	mockUserRepo.On("GetUserById", mock.Anything).Return(models.Users{
		Username: "test",
		Password: "test",
		Email:    "test@gmail.com",
	}, nil)

	response, err := usecase.Profile(uint(reqId))
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "test", response.Username)
	assert.Equal(t, "", response.Password)
	assert.Equal(t, "test@gmail.com", response.Email)
	mockUserRepo.AssertExpectations(t)
	mockBalanceRepo.AssertExpectations(t)
}
