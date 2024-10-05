package middleware

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password string) (string, error) {

	if len(password) < 8 {
		return "", errors.New("Password must be at least 8 characters long")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
