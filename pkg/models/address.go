package models

type Country struct {
	ID          int    `json:"id"`
	CountryName string `json:"country_name"`
}

type City struct {
	ID        int     `json:"id"`
	CityName  string  `json:"city_name"`
	Country   Country `json:"country"`
	CountryId int
}

type Address struct {
	ID       int  `json:"id"`
	PostCode int  `json:"postcode"`
	City     City `json:"city"`
	CityId   int
}
