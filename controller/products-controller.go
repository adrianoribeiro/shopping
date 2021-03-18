package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shopping/repo/models"
	"github.com/shopping/repo/services"
	"strconv"
)

func NewProduct(ctx *fiber.Ctx) error {

	var data map[string]interface{}
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	product := convertToProduct(data)
	if err := services.NewProduct(&product); err != nil {
		return err
	}

	return ctx.Status(201).JSON(product)
}

func GetProducts(ctx *fiber.Ctx) error {

	products := services.GetProducts()
	return ctx.JSON(products)
}

func GetProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	found, product := services.GetProduct(id)
	if !found {
		return nil //ctx.Status(204).SendString("No product found with this ID.")
	}

	return ctx.JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))
	if !services.Delete(id) {
		return nil //ctx.Status(204).SendString("Product delete successfully.")
	}
	return nil //ctx.Status(204).SendString("No product found with this ID.")
}

//Private functions
func convertToProduct(data map[string]interface{}) models.Product {
	return models.Product{
		Title:       data["title"].(string),
		Type:        data["type"].(string),
		Description: data["description"].(string),
		Filename:    data["filename"].(string),
		Height:      data["height"].(float64),
		Width:       data["width"].(float64),
		Price:       data["price"].(float64),
		Rating:      data["rating"].(float64),
	}
}
