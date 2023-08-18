package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	ans, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(ans), err
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}