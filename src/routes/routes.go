package routes

import (
	"go-shop-api/src/controllers/item"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	v1 := app.Group("v1")
	_ = v1.Group("api")
	admin := v1.Group("admin")
	admin.Get("/items", item.GetAllItems)
	admin.Post("/items", item.InsertItem)
}
