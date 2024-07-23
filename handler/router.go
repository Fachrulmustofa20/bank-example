package handler

import (
	"github.com/Fachrulmustofa20/bank-example.git/middleware"
	"github.com/Fachrulmustofa20/bank-example.git/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase    service.UsersUsecase
	balanceUsecase service.BalanceUsecase
	bankUsecase    service.BankUsecase
}

func NewUserHandler(r *gin.Engine,
	userUsecase service.UsersUsecase,
	balanceUsecase service.BalanceUsecase,
	bankUsecase service.BankUsecase,
) {
	handler := &Handler{
		userUsecase:    userUsecase,
		balanceUsecase: balanceUsecase,
		bankUsecase:    bankUsecase,
	}
	// test
	r.GET("/api/welcome", handler.Welcome)

	// users
	r.POST("/api/users/register", handler.Register)
	r.POST("/api/users/login", handler.Login)

	// banks
	r.POST("/api/bank/account", middleware.Authentication(), handler.CreateAccountBank)
	r.PUT("/api/bank/deposit", middleware.Authentication(), handler.AddDeposit)

	// balance
	r.GET("/api/balance", middleware.Authentication(), handler.GetBalance)
	r.POST("/api/balance/top-up", middleware.Authentication(), handler.TopUpBalance)
	r.POST("/api/balance/transfer", middleware.Authentication(), handler.TransferBalance)
	r.GET("/api/balance/mutation", middleware.Authentication(), handler.GetMutationBalance)
}
