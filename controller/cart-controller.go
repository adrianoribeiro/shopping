package controller

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shopping/repo/models"
	"github.com/shopping/repo/services"
)

type AddCartItemRequest struct {
	ProductId 	int `json:"productId"`
	Qty 		int `json:"qty"`
}

func Add(ctx *fiber.Ctx) error {

	token := ctx.Get("token", "")
	if len(token) == 0 {
		errors.New("no authenticate user")
	}

	var cartItem AddCartItemRequest
	if err := ctx.BodyParser(&cartItem); err != nil {
		return err
	}

	cartJson, _ := Rdx.HGet(token, "cart").Bytes()
	var cart models.Cart
	json.Unmarshal([]byte(cartJson), &cart)

	found, product := services.GetProduct(cartItem.ProductId)
	if found {
		cart.AddItem(models.CartItem{
			Product: product,
			Qty:     cartItem.Qty,
		})

		cartJson, _ := json.Marshal(cart)
		err := Rdx.HSet(token, "cart", cartJson).Err()
		if err != nil {
			panic(err)
		}
	}

	return ctx.Status(200).JSON(nil)
}

func GetCard(ctx *fiber.Ctx) error {
	token := ctx.Get("token", "")
	cartJson := Rdx.HGet(token, "cart").String()
	var cart models.Cart
	json.Unmarshal([]byte(cartJson), &cart)

	return ctx.JSON(cart)
}
