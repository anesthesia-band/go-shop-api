package middleware

import (
	"go-shop-api/src/common/middleware/auth/apikey"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func Init(app *fiber.App) {
	app.Use(keyauth.New(keyauth.Config{
		Next:      apikey.AdminAuthFilter,
		KeyLookup: "header:apiKey",
		Validator: apikey.ValidateAPIKey,
	}))
}
