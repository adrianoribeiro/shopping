package main

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/shopping/repo/controller"
	"github.com/shopping/repo/database"
	"github.com/shopping/repo/routes"
)

func init() {
	controller.Rdx = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func main(){

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
