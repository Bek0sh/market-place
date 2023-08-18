package controllers

import (
	"net/http"

	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/gin-gonic/gin"
)

type AddressController struct {
	repo irepository.AddressRepoInterface
}

func NewAddressController(repo irepository.AddressRepoInterface) *AddressController {
	return &AddressController{repo: repo}
}

func (cont AddressController) CreateAddress(ctx *gin.Context) {
	var address models.Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
		return
	}

	err := cont.repo.CreateAddress(&address)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "fail",
				"message": "failed to create address in database",
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func (cont *AddressController) UpdateAddress(ctx *gin.Context) {
	var address models.Address

	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)

		return
	}

	err := cont.repo.UpdateAddress(&address)

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
			"status":  "success",
			"message": "update",
		},
	)
}
