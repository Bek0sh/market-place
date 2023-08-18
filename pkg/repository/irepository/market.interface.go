package irepository

import "github.com/Bek0sh/market-place/pkg/models"

type MarketRepoInterface interface {
	GetProductById(id int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	DeleteProduct(id int) error
	DeleteAll() error
	UpdateProduct(product *models.Product) error
}
