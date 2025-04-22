package utils

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"regexp"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPasswordValidity(password string) bool {
	specialCharPattern := `[!@#$%^&*(),.?":{}|<>]`
	hasSpecialChar, _ := regexp.MatchString(specialCharPattern, password)
	return len(password) >= 8 && hasSpecialChar
}

func CheckValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
