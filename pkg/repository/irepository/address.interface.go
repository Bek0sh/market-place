package irepository

import "github.com/Bek0sh/market-place/pkg/models"

type AddressRepoInterface interface {
	GetAddressById(id int) (*models.Address, error)
	GetCityById(id int) (*models.City, error)
	GetCountryById(id int) (*models.Country, error)
	CreateAddress(address *models.Address) error
	DeleteAddress(id int) error
	UpdateAddress(address *models.Address) error
}