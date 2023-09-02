package service

import (
	"errors"

	"github.com/Bek0sh/market-place/pkg/config"
	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/Bek0sh/market-place/pkg/service/iservice"
	"github.com/Bek0sh/market-place/pkg/utils"
)

type authService struct {
	repo        irepository.AuthRepoInterface
	repoAddress irepository.AddressRepoInterface
}

var configs config.Config
var currentUserResp models.UserResponse

func NewAuthService(repo irepository.AuthRepoInterface, repoAddress irepository.AddressRepoInterface) iservice.AuthServiceInterface {
	configs, _ = config.LoadConfig(".")
	return &authService{repo: repo, repoAddress: repoAddress}
}

func (service *authService) Register(userInput *models.UserRegister) (int, error) {
	cityId, err := service.repoAddress.GetCityByName(userInput.Address.City.CityName)

	if err != nil {
		return 0, errors.New("failed to find city with this name in Database")
	}

	address := &models.Address{
		PostCode: userInput.Address.PostCode,
		CityId:   cityId.ID,
	}

	err = service.repoAddress.CreateAddress(address)

	if err != nil {
		return 0, errors.New("failed to create Address, error: " + err.Error())
	}

	hashedPassword, err := utils.HashPassword(userInput.Password)

	if err != nil || userInput.Password != userInput.ConfirmPassword {
		return 0, err
	}

	user := &models.User{
		Name:           userInput.Name,
		Surname:        userInput.Surname,
		Email:          userInput.Email,
		AddressId:      address.ID,
		HashedPassword: hashedPassword,
	}

	if userInput.Email == configs.AdminEmail && userInput.Password == configs.AdminPassword {
		user.UserType = "ADMIN"
	}

	id, err := service.repo.CreateUser(user)

	if err != nil {
		return 0, err
	}

	return id, nil

}

func (service *authService) SignIn(userInput *models.UserSignIn) (string, string, error) {
	user, err := service.repo.GetUserByEmail(userInput.Email)

	passwordMatch := utils.CheckPassword(userInput.Password, user.HashedPassword)

	if err != nil || passwordMatch != nil {
		return "", "", err
	}

	accessToken, err := utils.CreateToken(configs.AccessTokenExpiresIn, user.ID, configs.AccessTokenPrivateKey)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateToken(configs.RefreshTokenExpiresIn, user.ID, configs.RefreshTokenPrivateKey)

	if err != nil {
		return "", "", err
	}

	address, _ := service.repoAddress.GetAddressById(user.AddressId)

	currentUser := &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Address:  *address,
		Email:    user.Email,
		UserType: user.UserType,
	}

	currentUserResp = *currentUser

	return accessToken, refreshToken, nil

}

func (service authService) GetUserById(id int) (*models.UserResponse, error) {
	user, err := service.repo.GetUserById(id)
	if err != nil {
		return &models.UserResponse{}, err
	}

	address, err := service.repoAddress.GetAddressById(user.AddressId)

	if err != nil {
		return &models.UserResponse{}, err
	}

	userResponse := &models.UserResponse{
		Name:     user.Name,
		Surname:  user.Surname,
		UserType: user.UserType,
		Email:    user.Email,
		Address:  *address,
	}

	return userResponse, nil
}
