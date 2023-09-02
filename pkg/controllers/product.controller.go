package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/service/iservice"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service iservice.MarketServiceInterface
}

func NewProductController(service iservice.MarketServiceInterface) *ProductController {
	return &ProductController{service: service}
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

	err := cont.service.CreateProduct(&productInput)

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

	product, err := cont.service.GetProductById(id)

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

// func (cont *ProductController) UpdateProduct(ctx *gin.Context) {

// }

// func (cont *ProductController) GetProductByName(ctx *gin.Context) {

// }

func (cont *ProductController) DeleteProduct(ctx *gin.Context) {
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

	err = cont.service.DeleteProduct(id)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "fail",
				"message": "failed to delete product with this id",
			},
		)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":     "success",
			"deleted_id": id,
		},
	)
}

func (cont ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := cont.service.GetAllProducts()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "fail",
				"message": err.Error(),
			},
		)
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":   "success",
			"products": products,
		},
	)
}
