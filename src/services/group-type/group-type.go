package grouptype

import (
	"go-shop-api/src/common/response"
	grouptype "go-shop-api/src/repository/group-type"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	result, error := grouptype.GetAll()
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponseWithPayload(c, result)
}

func GetById(c *fiber.Ctx) error {
	id, error := extractItemIdFromRequest(c)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	result, error := grouptype.GetById(id)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponseWithPayload(c, result)
}

func Insert(c *fiber.Ctx) error {
	var data grouptype.InsertGroupTypeData
	error := c.BodyParser(&data)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	error = grouptype.Insert(data)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponse(c)
}

func UpdateById(c *fiber.Ctx) error {
	id, error := extractItemIdFromRequest(c)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	var data grouptype.InsertGroupTypeData
	error = c.BodyParser(&data)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	error = grouptype.UpdateById(id, data)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponse(c)
}

func DeleteById(c *fiber.Ctx) error {
	id, error := extractItemIdFromRequest(c)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}
	error = grouptype.DeleteById(id)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponse(c)
}

func extractItemIdFromRequest(c *fiber.Ctx) (uint, error) {
	id, error := strconv.ParseUint(c.Params("id"), 10, 32)
	return uint(id), error
}
