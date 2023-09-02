package routes

import (
	"github.com/Bek0sh/market-place/pkg/controllers"
	"github.com/Bek0sh/market-place/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(cont controllers.AuthController, router *gin.Engine) {

	router.POST("auth/register/", cont.Register)
	router.POST("auth/sign-in/", cont.SignIn)
	router.GET("/profile", middleware.CheckUser(), cont.Profile)
	router.GET("/logout", middleware.CheckUser(), cont.Logout)
}

func MarketRoutes(cont controllers.ProductController, router *gin.Engine) {
	router.POST("market/create", middleware.CheckUser(), cont.CreateProduct)
	router.GET("market/:id", cont.GetProductById)
	router.GET("market/", cont.GetAllProducts)
	router.DELETE("market/delete/:id", cont.DeleteProduct)
}
