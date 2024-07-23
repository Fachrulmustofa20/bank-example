package handler

import (
	"net/http"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
	"github.com/gin-gonic/gin"
)

func (handler Handler) GetBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	balance, err := handler.balanceUsecase.GetBalance(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}

func (handler Handler) TopUpBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	var topUp models.TopUpRequest

	if err := ctx.ShouldBindJSON(&topUp); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if topUp.CodeAccountBank == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please input your code in bank",
			"error":   "Code Bank is Empty",
		})
		return
	}

	if topUp.Amount < 10000 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Top up amount must be more than 10000",
			"error":   "Bad Request",
		})
		return
	}

	err := handler.balanceUsecase.TopUpBalance(topUp, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success top up balance. Check your balance",
	})
}

func (handler Handler) TransferBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	var transfer models.TransferBalance

	if err := ctx.ShouldBindJSON(&transfer); err != nil {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Transfer Your Balance",
	})
}

func (handler Handler) GetMutationBalance(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)

	history, err := handler.balanceUsecase.GetMutationBalance(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": history,
	})
}
