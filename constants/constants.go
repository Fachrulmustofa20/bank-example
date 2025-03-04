package constants

import (
	"errors"
)

var (
	AppPort = "APP_PORT"

	AppJSON = "application/json"

	ErrResponse         = errors.New("sign in to proceed")
	ErrExpired          = errors.New("token expired. please login again")
	ErrBalanceNotEnough = errors.New("your bank balance is not enough")
	ErrLogin            = errors.New("invalid email/password")
	ErrCodeBank         = errors.New("invalid code in bank")
)
