package controller

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shopping/repo/models"
	"github.com/shopping/repo/services"
	"math/rand"
)

func Payment(ctx *fiber.Ctx) error {

	token := ctx.Get("token", "")
	if len(token) == 0 {
		errors.New("no authenticate user")
	}

	payment := createPaymentObject(token)

	services.Publish(&payment)

	return ctx.Status(200).JSON(payment)
}

func createPaymentObject(token string) models.Payment {
	cart, shopUser := extractUserData(token)

	payment := models.Payment{
		User:       shopUser,
		CreditCard: generateCreditCard(dummyCreditCards),
		Amount:     cart.TotalPrice,
	}
	return payment
}

func extractUserData(token string) (models.Cart, models.ShopUser) {
	var cart models.Cart
	cartJson, err := Rdx.HGet(token, "cart").Bytes()
	json.Unmarshal(cartJson, &cart)
	if err != nil {
		errors.New("problem to serialize cart")
	}

	var shopUser models.ShopUser
	shopUserJson, err := Rdx.HGet(token, "shopUser").Bytes()
	json.Unmarshal(shopUserJson, &shopUser)
	if err != nil {
		errors.New("problem to serialize shop user")
	}
	return cart, shopUser
}

//
func generateCreditCard(dummyCreditCards []models.CreditCard) models.CreditCard {
	return dummyCreditCards[rand.Intn(5)]
}

var dummyCreditCards = []models.CreditCard{
	models.CreditCard{
		CardType:        "MasterCard",
		Name:            "Adriano Ribeiro",
		Number:          "3576093271612200",
		Cvv:             "100",
		ExpirationMonth: "10",
		ExpirationYear:  "2023",
	},
	models.CreditCard{
		CardType:        "MasterCard",
		Name:            "Abigail Abada",
		Number:          "3576093271612200",
		Cvv:             "101",
		ExpirationMonth: "09",
		ExpirationYear:  "2022",
	},
	models.CreditCard{
		CardType:        "MasterCard",
		Name:            "Benedita Bin",
		Number:          "3576093271613200",
		Cvv:             "102",
		ExpirationMonth: "01",
		ExpirationYear:  "2024",
	},
	models.CreditCard{
		CardType:        "MasterCard",
		Name:            "Carlotta Caldas",
		Number:          "3576093271614200",
		Cvv:             "100",
		ExpirationMonth: "04",
		ExpirationYear:  "2022",
	},
	models.CreditCard{
		CardType:        "MasterCard",
		Name:            "Demetria Drum",
		Number:          "3576093271616200",
		Cvv:             "100",
		ExpirationMonth: "10",
		ExpirationYear:  "2023",
	},
}
