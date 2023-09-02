package controllers

import (
	"net/http"

	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/service/iservice"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service iservice.AuthServiceInterface
}

func NewAuthController(service iservice.AuthServiceInterface) *AuthController {
	return &AuthController{service: service}
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

	id, err := cont.service.Register(&userInput)

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

	accessToken, refreshToken, err := cont.service.SignIn(&userInput)

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
	userId := ctx.MustGet("user_id")
	currentUser, err := cont.service.GetUserById(int(userId.(float64)))

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

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":       "success",
			"current_user": currentUser,
		},
	)
}
