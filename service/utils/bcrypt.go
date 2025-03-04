package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pass string) string {
	salt := 8
	password := []byte(pass)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
