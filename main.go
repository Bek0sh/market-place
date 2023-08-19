package main

import (
	"log"

	"github.com/Bek0sh/market-place/pkg/config"
	"github.com/Bek0sh/market-place/pkg/controllers"
	"github.com/Bek0sh/market-place/pkg/db"
	"github.com/Bek0sh/market-place/pkg/repository"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/Bek0sh/market-place/pkg/routes"
	"github.com/gin-gonic/gin"
)

var authRepo irepository.AuthRepoInterface
var addressRepo irepository.AddressRepoInterface
var marketRepo irepository.MarketRepoInterface

var authCont controllers.AuthController
var addressCont controllers.AddressController

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

	authCont = *controllers.NewAuthController(authRepo, addressRepo)
	addressCont = *controllers.NewAddressController(addressRepo)
}

func main() {
	router := gin.Default()

	routes.AuthRoutes(authCont, router)

	panic(router.Run(":8080"))
}
