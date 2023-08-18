package routes

import (
	"github.com/Bek0sh/market-place/pkg/controllers"
	"github.com/Bek0sh/market-place/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(cont controllers.AuthController, router *gin.Engine) {

	router.POST("auth/register/", cont.Register)
	router.POST("auth/sign-in/", cont.SignIn)
	router.GET("user/profile", middleware.CheckUser(), cont.Profile)
	router.GET("user/logout", middleware.CheckUser(), cont.Logout)
}