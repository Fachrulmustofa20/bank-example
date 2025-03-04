package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Fachrulmustofa20/bank-example.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (cfg *Config) InitPostgres() error {
	var (
		DBHost    = os.Getenv("DB_HOST")
		DBUser    = os.Getenv("DB_USER")
		DBPwd     = os.Getenv("DB_PWD")
		DBName    = os.Getenv("DB_NAME")
		DBPort    = os.Getenv("DB_PORT")
		DBSSLMode = os.Getenv("DB_SSL_MODE")
	)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		DBHost, DBUser, DBPwd, DBName, DBPort, DBSSLMode)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect to database, error: %+v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("[ERR]: error while connect to db: %+v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 15)

	cfg.db = db

	fmt.Println("Success connect to database")
	if err := db.AutoMigrate(&models.Users{},
		&models.Balance{},
		&models.BalanceHistory{},
		&models.Bank{},
		&models.BankBalanceHistory{}); err != nil {
		log.Fatalln(err)
	}
	return nil
}
