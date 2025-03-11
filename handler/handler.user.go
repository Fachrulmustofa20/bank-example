package handler

import (
	"net/http"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"github.com/Fachrulmustofa20/bank-example.git/service/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Register ... Create User
// @Summary Create new user
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body models.RegisterRequest true "User Data"
// @Success 201 {object} object
// @Failure 422,500 {object} object
// @Router /users/register [post]
func (handler Handler) Register(ctx *gin.Context) {
	var registerReq models.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		log.Error("error binding json: ", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(registerReq)
	if err != nil || !valid {
		log.Warn("error validate struct: ", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A validation error occurred. Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	err = handler.userUsecase.Register(registerReq)
	if err != nil {
		log.Error("error register user: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Created Account!",
	})
}

// Login ... Login User
// @Summary Login user
// @Description Login user
// @Tags Users
// @Accept json
// @Param user body models.LoginRequest true "User Data"
// @Success 200 {object} object
// @Failure 422,500 {object} object
// @Router /users/login [post]
func (handler Handler) Login(ctx *gin.Context) {
	var loginReq models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		log.Error("error binding json: ", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	valid, err := govalidator.ValidateStruct(loginReq)
	if err != nil || !valid {
		log.Warn("error validate struct: ", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "A validation error occurred. Please check the form and try again.",
			"error":   err.Error(),
		})
		return
	}

	password := loginReq.Password
	token, err := handler.userUsecase.Login(loginReq.Email, password)
	if err != nil {
		log.Error("error login user: ", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"access_token": token,
	})
}

// Profile ... Get Profile Users
// @Summary Get Profile Users
// @Description Get Profile Users
// @Tags Users
// @Accept json
// @Success 200 {object} object
// @Security JWT
// @Failure 422,500 {object} object
// @Router /users/profile [get]
func (handler Handler) Profile(ctx *gin.Context) {
	userId := utils.GetUserIdJWT(ctx)
	users, err := handler.userUsecase.Profile(userId)
	if err != nil {
		log.Error("error get profile user: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})
}
