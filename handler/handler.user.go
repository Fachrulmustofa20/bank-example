package handler

import (
	"net/http"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (handler Handler) Register(ctx *gin.Context) {
	var users models.Users
	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(users)
	if err != nil || !valid {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A validation error occurred. Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	err = handler.userUsecase.Register(users)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status Bad Request",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Created Account!",
	})
}

func (handler Handler) Login(ctx *gin.Context) {
	var users models.Users
	if err := ctx.ShouldBindJSON(&users); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	if users.Email == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please input your email!",
		})
		return
	}
	if users.Password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please input your password!",
		})
		return
	}

	password := users.Password
	token, err := handler.userUsecase.Login(users.Email, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
