package main

import (
	"go-shop-api/src/common/database"
	"go-shop-api/src/common/middleware"
	"go-shop-api/src/controller"
	"go-shop-api/src/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.Init()
	app := fiber.New()
	middleware.Init(app)
	repository.Init(db)
	controller.Init(app)

	app.Listen(":3000")
}
