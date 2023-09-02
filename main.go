package main

import (
	"log"

	"github.com/Bek0sh/market-place/pkg/config"
	"github.com/Bek0sh/market-place/pkg/controllers"
	"github.com/Bek0sh/market-place/pkg/db"
	"github.com/Bek0sh/market-place/pkg/repository"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/Bek0sh/market-place/pkg/routes"
	"github.com/Bek0sh/market-place/pkg/service"
	"github.com/Bek0sh/market-place/pkg/service/iservice"
	"github.com/gin-gonic/gin"
)

var authRepo irepository.AuthRepoInterface
var addressRepo irepository.AddressRepoInterface
var marketRepo irepository.MarketRepoInterface

var authService iservice.AuthServiceInterface
var marketService iservice.MarketServiceInterface

var authCont controllers.AuthController
var addressCont controllers.AddressController
var marketCont controllers.ProductController

func init() {
	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Println("failed to load config")
		log.Fatal(err.Error())
	}
	db.Connect(cfg)

	authRepo = repository.NewauthRepository(db.DB)
	addressRepo = repository.NewAddressRepository(db.DB)
	marketRepo = repository.NewMarketRepository(db.DB)

	authService = service.NewAuthService(authRepo, addressRepo)
	marketService = service.NewMarketService(marketRepo)

	authCont = *controllers.NewAuthController(authService)
	marketCont = *controllers.NewProductController(marketService)
	addressCont = *controllers.NewAddressController(addressRepo)
}

func main() {
	router := gin.Default()

	routes.AuthRoutes(authCont, router)
	routes.MarketRoutes(marketCont, router)

	panic(router.Run(":8080"))
}
