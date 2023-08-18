package repository

import (
	"github.com/Bek0sh/market-place/pkg/models"
	"github.com/Bek0sh/market-place/pkg/repository/irepository"
	"gorm.io/gorm"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) irepository.AddressRepoInterface {
	return &addressRepo{db: db}
}

func (repo *addressRepo) GetAddressById(id int) (*models.Address, error) {
	var address models.Address

	if err := repo.db.First(&address, "id=?", id).Error; err != nil {
		return &models.Address{}, err
	}

	return &address, nil
}

func (repo *addressRepo) GetCityById(id int) (*models.City, error) {
	var city models.City

	if err := repo.db.First(&city, "id=?", id).Error; err != nil {
		return &models.City{}, err
	}

	return &city, nil
}

func (repo *addressRepo) GetCountryById(id int) (*models.Country, error) {
	var country models.Country

	if err := repo.db.First(&country, "id=?", id).Error; err != nil {
		return &models.Country{}, err
	}

	return &country, nil
}

func (repo *addressRepo) CreateAddress(address *models.Address) error {
	return repo.db.Create(&address).Error
}

func (repo *addressRepo) DeleteAddress(id int) error {
	if err := repo.db.Where("id=?", id).Delete(&models.Address{}).Error; err != nil {
		return err
	}

	return nil
}

func (repo *addressRepo) UpdateAddress(address *models.Address) error {
	if err := repo.db.Model(&models.Product{}).Where("id=?", address.ID).Updates(&models.Address{
		PostCode: address.PostCode,
		CityId:   address.CityId,
	}).Error; err != nil {
		return err
	}

	return nil
}
