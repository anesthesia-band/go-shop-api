package controller

import (
	"go-shop-api/src/services/item"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	v1 := app.Group("v1")
	initAdminControllers(v1)
	initClientControllers(v1)
}

func initAdminControllers(app fiber.Router) {
	admin := app.Group("admin/api")
	admin.Get("/items", item.GetAllItems)
	admin.Post("/items", item.InsertItem)
}

func initClientControllers(app fiber.Router) {
	_ = app.Group("api")
}
