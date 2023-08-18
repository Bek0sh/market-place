package irepository

import "github.com/Bek0sh/market-place/pkg/models"

type AuthRepoInterface interface {
	CreateUser(user *models.User) (int, error)
	SignIn(user *models.UserSignIn) error
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(id int) (int, error)
}
