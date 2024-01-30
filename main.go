package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// host=localhost user=gorm password=gorm dbname=gorm port=9920
	_, err := gorm.Open(postgres.Open("host=db user=root password=root dbname=shop_db port=5432"), &gorm.Config{})

	if (err != nil) {
		panic("db not connected")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}