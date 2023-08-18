package models

type User struct {
	ID             int
	Name           string
	UserType       string
	Surname        string
	Email          string
	AddressId      int
	HashedPassword string
}

type UserResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	UserType  string  `json:"user_type"`
	Surname   string  `json:"surname"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
	AddressId int
}

type UserRegister struct {
	Name            string  `json:"name"`
	Surname         string  `json:"surname"`
	Email           string  `json:"email"`
	Password        string  `json:"password"`
	Address         Address `json:"address"`
	AddressId       int
	ConfirmPassword string `json:"confirm_password"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
