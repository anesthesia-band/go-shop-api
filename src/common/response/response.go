package response

import (
	"github.com/gofiber/fiber/v2"
)

type data any

func SuccessResponse(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "Ok",
	})
}

func SuccessResponseWithPayload(c *fiber.Ctx, data data) error {
	return c.JSON(fiber.Map{
		"status":  "Ok",
		"payload": data,
	})
}

func BadRequestResponse(c *fiber.Ctx, error error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "Error",
		"payload": error.Error(),
	})
}

func InternalServerResponse(c *fiber.Ctx, error error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "Error",
		"payload": error.Error(),
	})
}

func IAmATeapotResponse(c *fiber.Ctx, error error) error {
	return c.Status(fiber.StatusTeapot).JSON(fiber.Map{
		"status": "Error",
	})
}
