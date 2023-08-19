package controllers

import (
	"net/http"

	"github.com/Bek0sh/market-place/pkg/config"
	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/Bek0sh/market-place/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	repo        irepository.AuthRepoInterface
	addressRepo irepository.AddressRepoInterface
}

var configs config.Config

func NewAuthController(repo irepository.AuthRepoInterface, addressRepo irepository.AddressRepoInterface) *AuthController {
	configs, _ = config.LoadConfig(".")
	return &AuthController{repo: repo, addressRepo: addressRepo}
}

func (cont AuthController) Register(ctx *gin.Context) {
	var userInput models.UserRegister

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	cityId, err := cont.addressRepo.GetCityByName(userInput.Address.City.CityName)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "fail",
				"message": "failed to find city with this name in Database",
			},
		)
		return
	}

	address := &models.Address{
		PostCode: userInput.Address.PostCode,
		CityId:   cityId.ID,
	}

	err = cont.addressRepo.CreateAddress(address)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "fail",
				"message": "failed to find city with this name in Database",
			},
		)
		return
	}

	hashedPassword, err := utils.HashPassword(userInput.Password)

	if err != nil || userInput.Password != userInput.ConfirmPassword {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
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

	id, err := cont.repo.CreateUser(user)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":          "success",
			"created_user_id": id,
		},
	)
}

func (cont AuthController) SignIn(ctx *gin.Context) {
	var userInput models.UserSignIn

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	user, err := cont.repo.GetUserByEmail(userInput.Email)

	passwordMatch := utils.CheckPassword(userInput.Password, user.HashedPassword)

	if err != nil || passwordMatch != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	accessToken, err := utils.CreateToken(configs.AccessTokenExpiresIn, user.ID, configs.AccessTokenPrivateKey)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	refreshToken, err := utils.CreateToken(configs.RefreshTokenExpiresIn, user.ID, configs.RefreshTokenPrivateKey)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	ctx.SetCookie("access_token", accessToken, 20*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, 100*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", 20*60, "/", "localhost", false, false)

	address, _ := cont.addressRepo.GetAddressById(user.AddressId)

	currentUser := &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Address:  *address,
		Email:    user.Email,
		UserType: user.UserType,
	}

	ctx.Set("current_user", currentUser)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":       "success",
			"access_token": accessToken,
		},
	)
}

func (cont AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "api/v1", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "api/v1", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "api/v1", false, false)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "logged out",
		},
	)
}

func (cont AuthController) Profile(ctx *gin.Context) {
	currentUser := ctx.MustGet("current_user")

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":       "success",
			"current_user": currentUser,
		},
	)
}
