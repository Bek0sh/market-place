package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	repo irepository.MarketRepoInterface
}

func NewProductController(repo irepository.MarketRepoInterface) *ProductController {
	return &ProductController{repo: repo}
}

func (cont *ProductController) CreateProduct(ctx *gin.Context) {
	var productInput models.Product

	if err := ctx.ShouldBindJSON(&productInput); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": "invalid input",
			},
		)
		return
	}

	err := cont.repo.CreateProduct(&productInput)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": "failed to create product",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":     "success",
			"product_id": productInput.ID,
		},
	)
}

func (cont *ProductController) GetProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "fail",
				"message": "failed to convert id",
			},
		)
	}

	product, err := cont.repo.GetProductById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "fail",
				"message": "failed to find product with this id",
			},
		)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"product": product,
		},
	)

}

func (cont *ProductController) UpdateProduct(ctx *gin.Context) {

}

func (cont *ProductController) GetProductByName(ctx *gin.Context) {

}

func (cont *ProductController) DeleteProduct(ctx *gin.Context) {

}
