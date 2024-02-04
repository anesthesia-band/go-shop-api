package main

import (
	"go-shop-api/src/common/database"
	"go-shop-api/src/common/middleware"
	logger "go-shop-api/src/common/server-logger"
	"go-shop-api/src/controller"
	"go-shop-api/src/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	db := database.Init()

	app := fiber.New()
	app.Use(requestid.New())
	logger.Init(app)

	middleware.Init(app)
	repository.Init(db)
	controller.Init(app)

	app.Listen(":3000")
}
