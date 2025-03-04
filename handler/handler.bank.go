package handler

import (
	"net/http"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// CreateAccountBank ... Create Account Bank
// @Summary Create new Account Bank
// @Description Create new Account Bank
// @Tags Bank
// @Accept json
// @Param user body models.CreateAccountBankRequest true "Account Bank"
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /bank/account [post]
func (handler Handler) CreateAccountBank(ctx *gin.Context) {
	var request models.CreateAccountBankRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(request)
	if err != nil || !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A validation error occurred. Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if request.Balance < 500000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
			"error":   "Balance must be above 500000!",
		})
		return
	}

	userId := utils.GetUserIdJWT(ctx)
	err = handler.bankUsecase.CreateAccountBank(models.Bank{
		Balance:        request.Balance,
		Code:           request.Code,
		BalanceAchieve: request.BalanceAchieve,
		UserId:         userId,
		Enable:         true,
	})
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
			"balance":         request.Balance,
			"balance_achieve": request.BalanceAchieve,
			"code":            request.Code,
		},
	})
}

// AddDeposit ... Add Deposit
// @Summary Add Deposit
// @Description Add Deposit
// @Tags Bank
// @Accept json
// @Param user body models.AddDepositRequest true "Add Deposit"
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /bank/deposit [put]
func (handler Handler) AddDeposit(ctx *gin.Context) {
	var request models.AddDepositRequest
	userId := utils.GetUserIdJWT(ctx)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if request.Balance < 100000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "please input minimum deposit 100000",
		})
		return
	}

	if request.Code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   "please input your code to continue transaction",
		})
		return
	}

	err := handler.bankUsecase.AddDeposit(models.Bank{
		Balance:        request.Balance,
		BalanceAchieve: request.BalanceAchieve,
		Code:           request.Code,
		Enable:         true,
		UserId:         userId,
	})
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
			"balance": request.Balance,
			"user_id": userId,
		},
	})
}
