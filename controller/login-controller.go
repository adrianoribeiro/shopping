package controller

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/shopping/repo/models"
	"strconv"
	"time"
)

var redisCtx = context.Background()
var Rdx *redis.Client
func Login(ctx *fiber.Ctx) error {

	var shopUser models.ShopUser
	ctx.BodyParser(&shopUser)

	//Add token.
	//This token is used to validate the requests.
	//This approach is to avoid implement security (just to keep simple)
	shopUser.Token = strconv.FormatInt(time.Now().UnixNano(), 10)

	shopUserJson, _ := json.Marshal(shopUser)
	Rdx.HSet(shopUser.Token, "shopUser", shopUserJson)

	return ctx.Status(200).JSON(shopUser.Token)
}


