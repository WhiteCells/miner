package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to encrypt password")
	}
	return string(hashPassword), nil
}

func ValidPassword(userPassword string, inPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inPassword)) == nil
}
