package iservice

import "github.com/Bek0sh/market-place/pkg/models"

type AuthServiceInterface interface {
	Register(*models.UserRegister) (int, error)
	SignIn(*models.UserSignIn) (string, string, error)
	GetUserById(int) (*models.UserResponse, error)
}
