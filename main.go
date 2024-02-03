package main

import (
	"go-shop-api/src/common/middleware"
	"go-shop-api/src/database"
	"go-shop-api/src/routes"
	"go-shop-api/src/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.Init()
	app := fiber.New()
	middleware.Init(app)
	services.Init(db)
	routes.Init(app)

	app.Listen(":3000")
}
