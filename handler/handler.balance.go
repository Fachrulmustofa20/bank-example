package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
	"github.com/gin-gonic/gin"
)

// GetBalance ... Get Balance
// @Summary Get Balance
// @Description Get Balance
// @Tags Balance
// @Accept json
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /balance [get]
func (handler Handler) GetBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	balance, err := handler.balanceUsecase.GetBalance(userId)
	if err != nil {
		log.Error("error get balance: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"balance": balance,
	})
}

// TopUpBalance ... Top Up Balance
// @Summary Top Up Balance
// @Description Top Up Balance
// @Tags Balance
// @Accept json
// @Param topUp body models.TopUpRequest true "Top Up Data"
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /balance/top-up [post]
func (handler Handler) TopUpBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	var topUp models.TopUpRequest

	if err := ctx.ShouldBindJSON(&topUp); err != nil {
		log.Error("error binding json: ", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if topUp.CodeAccountBank == "" {
		log.Info("error code bank is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input your code in bank",
			"error":   "Code Bank is Empty",
		})
		return
	}

	if topUp.Amount < 10000 {
		log.Info("error top up amount must be more than 10000")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Top up amount must be more than 10000",
			"error":   "Bad Request",
		})
		return
	}

	err := handler.balanceUsecase.TopUpBalance(topUp, userId)
	if err != nil {
		log.Error("error top up balance: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success top up balance. Check your balance",
	})
}

// TransferBalance ... Transfer Balance
// @Summary Transfer Balance
// @Description Transfer Balance
// @Tags Balance
// @Accept json
// @Param transfer body models.TransferBalance true "Transfer Balance"
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /balance/transfer [post]
func (handler Handler) TransferBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	var transfer models.TransferBalance

	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		log.Error("error bind json: ", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if transfer.EmailRecipient == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input email recepient",
			"error":   "Email Recepient is Empty",
		})
		return
	}

	if transfer.Amount < 10000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Transfer amount must be more than 10000",
			"error":   "Bad Request",
		})
		return
	}

	err := handler.balanceUsecase.TransferBalance(transfer, userId)
	if err != nil {
		log.Error("error transfer balance: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad Request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Transfer Your Balance",
	})
}

// GetMutationBalance ... Get Mutation Balance
// @Summary Get Mutation Balance
// @Description Get Mutation Balance
// @Tags Balance
// @Accept json
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /balance/mutation [get]
func (handler Handler) GetMutationBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	history, err := handler.balanceUsecase.GetMutationBalance(userId)
	if err != nil {
		log.Error("error get mutation balance: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    history,
	})
}
