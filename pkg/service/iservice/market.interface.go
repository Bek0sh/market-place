package iservice

import "github.com/Bek0sh/market-place/pkg/models"

type MarketServiceInterface interface {
	DeleteProduct(id int) error
	UpdateProduct(*models.Product) error
	CreateProduct(*models.Product) error
	GetProductById(int) (*models.Product, error)
	GetAllProducts() ([]models.Product, error)
}
