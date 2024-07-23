package handler

import (
	"net/http"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (handler Handler) CreateAccountBank(ctx *gin.Context) {
	var bank models.Bank
	userId := utils.GetUserIdJWT(ctx)
	bank.UserId = userId
	bank.Enable = true

	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(bank)
	if err != nil || !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A validation error occurred. Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if bank.Balance < 500000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
			"error":   "Balance must be above 500000!",
		})
		return
	}

	err = handler.bankUsecase.CreateAccountBank(bank)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Created Account Bank!",
		"data": map[string]interface{}{
			"balance":         bank.Balance,
			"balance_achieve": bank.BalanceAchieve,
			"code":            bank.Code,
		},
	})
}

func (handler Handler) AddDeposit(ctx *gin.Context) {
	var bank models.Bank
	userId := utils.GetUserIdJWT(ctx)
	bank.UserId = userId

	if err := ctx.ShouldBindJSON(&bank); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if bank.Balance < 100000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "please input minimum deposit 100000",
		})
		return
	}

	if bank.Code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "please input your code to continue transaction",
		})
		return
	}

	err := handler.bankUsecase.AddDeposit(bank)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Deposit!",
		"data": map[string]interface{}{
			"balance": bank.Balance,
			"user_id": userId,
		},
	})
}
