package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func ValidatePassword(hashPassword, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Password))
}
