package models

type Product struct {
	ID         int          `json:"-"`
	Name       string       `json:"name"`
	Price      int          `json:"price"`
	User       UserResponse `json:"created_by"`
	UserId     int
	Address    Address `json:"address"`
	AddressId  int
	Category   Category `json:"category"`
	CategoryId int
}

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}
