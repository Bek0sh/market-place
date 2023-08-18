package repository

import (
	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/Bek0sh/market-place/pkg/utils"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewauthRepository(db *gorm.DB) irepository.AuthRepoInterface {
	return &authRepo{db: db}
}

func (repo *authRepo) CreateUser(user *models.User) (int, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (repo *authRepo) SignIn(user *models.UserSignIn) error {
	var password string

	if err := repo.db.First(&password, "email=?", user.Email).Error; err != nil {
		return err
	}

	err := utils.CheckPassword(user.Password, password)

	return err
}

func (repo *authRepo) GetUserById(id int) (*models.User, error) {
	var user models.User

	if err := repo.db.Where("id=?", id).Find(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (repo *authRepo) DeleteUser(id int) (int, error) {

	if err := repo.db.Where("id=?", id).Delete(&models.User{}).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *authRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := repo.db.Where("email=?", email).Find(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
