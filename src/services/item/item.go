package item

import (
	"go-shop-api/src/repository/item"

	"github.com/gofiber/fiber/v2"
)

func InsertItem(c *fiber.Ctx) error {
	var data item.InsertItemData
	error := c.BodyParser(&data)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"payload": error.Error(),
		})
	}

	error = item.InsertItem(data)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"payload": error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "Ok",
	})
}

func GetAllItems(c *fiber.Ctx) error {
	items, error := item.GetItems()
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"payload": error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "Ok",
		"payload": items,
	})
}
