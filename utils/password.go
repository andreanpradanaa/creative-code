package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashPassword), err
}

func VerifyPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
