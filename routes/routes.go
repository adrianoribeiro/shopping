package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shopping/repo/controller"
)

func Setup(app *fiber.App) {
	//Login
	app.Post("/login", controller.Login)

	//Products
	app.Post("/products", controller.NewProduct)
	app.Get("/products", controller.GetProducts)
	app.Get("/products/:id", controller.GetProduct)
	app.Delete("/products/:id", controller.DeleteProduct)

	//Cart
	app.Get("/cart", controller.GetCard)
	app.Post("/cart/add", controller.Add)

	//Payment
	app.Post("/payment", controller.Payment)
}
