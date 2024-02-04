package groupitem

import (
	"go-shop-api/src/repository"
	"go-shop-api/src/repository/models"
)

type InsertGroupItemData struct {
	GroupId  uint
	ItemId   uint
	Quantity uint
}

type UpdateGroupItemData struct {
	GroupId  uint
	ItemId   uint
	Quantity uint
	Active   bool
}

func GetByGroupId(groupId uint) ([]models.GroupItem, error) {
	var groupItems []models.GroupItem
	result := repository.DB.Where("group_id = ?", groupId).Find(&groupItems)
	return groupItems, result.Error
}

func Insert(data InsertGroupItemData) error {
	result := repository.DB.Create(&models.GroupItem{
		GroupId:  data.GroupId,
		ItemId:   data.ItemId,
		Quantity: data.Quantity,
	})
	return result.Error
}

func DeleteByGroupIdAndItempId(groupId uint, itemId uint) error {
	result := repository.DB.Where("group_id = ?", groupId).Where("item_id = ?", itemId).Delete(&models.ItemType{})
	return result.Error
}
