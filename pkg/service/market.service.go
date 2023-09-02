package service

import (
	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"github.com/Bek0sh/market-place/pkg/service/iservice"
)

type marketService struct {
	repo irepository.MarketRepoInterface
}

func NewMarketService(repo irepository.MarketRepoInterface) iservice.MarketServiceInterface {
	return &marketService{repo: repo}
}

func (service *marketService) DeleteProduct(id int) error {
	product, err := service.repo.GetProductById(id)

	if err != nil || product.UserId != currentUserResp.ID {
		return err
	}

	err = service.repo.DeleteProduct(id)

	return err
}

func (service *marketService) UpdateProduct(product *models.Product) error {
	product, err := service.repo.GetProductById(0)

	if err != nil || product.UserId != currentUserResp.ID {
		return err
	}
	return nil
}

func (service *marketService) CreateProduct(product *models.Product) error {
	product.AddressId = currentUserResp.AddressId
	product.UserId = currentUserResp.ID
	catogoryId, err := service.repo.GetCategoryByName(product.Category.CategoryName)
	if err != nil {
		return err
	}
	product.CategoryId = catogoryId

	return service.CreateProduct(product)
}

func (service *marketService) GetProductById(id int) (*models.Product, error) {
	return service.repo.GetProductById(id)
}

func (service *marketService) GetAllProducts() ([]models.Product, error) {
	return service.GetAllProducts()
}
