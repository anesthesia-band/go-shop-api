package groupitem

import (
	"go-shop-api/src/common/response"
	groupitem "go-shop-api/src/repository/group-item"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetByGroupId(c *fiber.Ctx) error {
	id, error := extractItemIdFromRequest(c)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	result, error := groupitem.GetByGroupId(id)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponseWithPayload(c, result)
}

func Insert(c *fiber.Ctx) error {
	var data groupitem.InsertGroupItemData
	error := c.BodyParser(&data)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	error = groupitem.Insert(data)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponse(c)
}

func DeleteByGroupIdAndItempId(c *fiber.Ctx) error {
	groupId, error := extractItemIdFromRequest(c)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	itemId, error := extractItemIdFromRequest(c)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	error = groupitem.DeleteByGroupIdAndItempId(groupId, itemId)
	if error != nil {
		return response.BadRequestResponse(c, error)
	}

	return response.SuccessResponse(c)
}

func extractItemIdFromRequest(c *fiber.Ctx) (uint, error) {
	id, error := strconv.ParseUint(c.Queries()["itemId"], 10, 32)
	return uint(id), error
}

func extractGroupIdFromRequest(c *fiber.Ctx) (uint, error) {
	id, error := strconv.ParseUint(c.Queries()["groupId"], 10, 32)
	return uint(id), error
}
