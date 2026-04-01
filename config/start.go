package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Fachrulmustofa20/bank-example.git/handler"
	"github.com/Fachrulmustofa20/bank-example.git/service/repository/postgres"
	"github.com/Fachrulmustofa20/bank-example.git/service/usecase"
	"github.com/gin-gonic/gin"
)

func (cfg *Config) Start() error {
	port := os.Getenv("APP_PORT")
	appPort := fmt.Sprintf(":%s", port)

	r := gin.Default()
	if strings.ToLower(os.Getenv("ENV")) == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// init repo
	userRepo := postgres.NewUserRepository(cfg.db)
	balanceRepo := postgres.NewBalanceRepository(cfg.db)
	bankRepo := postgres.NewBankRepository(cfg.db)

	// init usecase
	userUsecase := usecase.NewUsersUsecase(userRepo, balanceRepo)
	balanceUsecase := usecase.NewBalanceUsecase(userRepo, balanceRepo, bankRepo)
	bankUsecase := usecase.NewBankUsecase(bankRepo, userRepo)

	handler.NewUserHandler(r, userUsecase, balanceUsecase, bankUsecase)

	srv := &http.Server{
		Addr:    appPort,
		Handler: r,
	}

	// channel untuk OS signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)

	// run server di goroutine
	go func() {
		log.Println("HTTP server running on", appPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("HTTP server error: %w", err)
		}
	}()

	// wait event
	select {
	case sig := <-sigs:
		log.Println("Shutting down service due to signal:", sig.String())
	case err := <-errChan:
		log.Println("Service exited with error:", err)
	}

	// graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Failed to shutdown HTTP server:", err)
		return err
	}

	log.Println("Service stopped gracefully")
	return nil
}
