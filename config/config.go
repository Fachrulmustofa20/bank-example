package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Fachrulmustofa20/bank-example.git/handler"
	"github.com/Fachrulmustofa20/bank-example.git/service/repository/postgres"
	"github.com/Fachrulmustofa20/bank-example.git/service/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB
}

func Init() Config {
	var cfg Config
	err := cfg.InitPostgres()
	if err != nil {
		log.Panic()
	}

	fmt.Println("Server is running ..")

	return cfg
}

func (cfg *Config) Start() error {
	port := os.Getenv("APP_PORT")
	appPort := fmt.Sprintf(":%s", port)
	r := gin.Default()

	// init repo
	userRepo := postgres.NewUserRepository(cfg.db)
	balanceRepo := postgres.NewBalanceRepository(cfg.db)
	bankRepo := postgres.NewBankRepository(cfg.db)
	// init usecase
	userUsecase := usecase.NewUsersUsecase(userRepo, balanceRepo)
	balanceUsecase := usecase.NewBalanceUsecase(userRepo, balanceRepo, bankRepo)
	bankUsecase := usecase.NewBankUsecase(bankRepo, userRepo)
	handler.NewUserHandler(r, userUsecase, balanceUsecase, bankUsecase)

	err := r.Run(appPort)
	if err != nil {
		fmt.Printf("[ERR] Start service %+v", err)
	}
	return nil
}
