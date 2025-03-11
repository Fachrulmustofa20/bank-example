package config

import (
	"fmt"
	"os"

	"github.com/Fachrulmustofa20/bank-example.git/handler"
	"github.com/Fachrulmustofa20/bank-example.git/service/repository/postgres"
	"github.com/Fachrulmustofa20/bank-example.git/service/usecase"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB
}

func Init() Config {
	var cfg Config
	err := cfg.initPostgres()
	if err != nil {
		log.Panic()
	}

	initLogger()
	log.Info("Server is running ..")

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
		log.Error("[ERR] Start service ", err)
	}
	return nil
}
