package models

type Payment struct {
	User ShopUser
	Amount float64
	CreditCard CreditCard `json:"credit_card"`
}

type CreditCard struct {
	CardType string `json:"card_type"`
	Name string
	Number string
	Cvv	string
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear string `json:"expiration_year"`
}
