package repository

import (
	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"gorm.io/gorm"
)

type marketRepo struct {
	db *gorm.DB
}

func NewMarketRepository(db *gorm.DB) irepository.MarketRepoInterface {
	return &marketRepo{db: db}
}

func (repo marketRepo) GetProductById(id int) (*models.Product, error) {
	var product models.Product

	if err := repo.db.First(&product, "id=?", id).Error; err != nil {
		return &models.Product{}, err
	}

	return &product, nil
}

func (repo marketRepo) CreateProduct(product *models.Product) error {
	if err := repo.db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func (repo marketRepo) DeleteProduct(id int) error {
	if err := repo.db.Where("id=?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (repo marketRepo) DeleteAll() error {
	return repo.db.Delete(&models.Product{}).Error
}

func (repo marketRepo) UpdateProduct(product *models.Product) error {

	if err := repo.db.Model(&models.Product{}).Where("id=?", product.ID).Updates(&models.Product{
		Name:    product.Name,
		Price:   product.Price,
		Address: product.Address,
	}).Error; err != nil {
		return err
	}

	return nil
}
