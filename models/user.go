package models

type Address struct {
	Street 		string
	PostalCode 	string
	City 		string
	CountryCode string
	Country 	string
}

type ShopUser struct {
	Name string
	Address Address
	Token string
}

