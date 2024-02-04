package item

import (
	"go-shop-api/src/repository"
	"go-shop-api/src/repository/models"
)

type InsertItemData struct {
	Name       string
	ItemTypeId uint `gorm:"index"`
	Data       string
	Price      string
}

func GetItems() ([]models.Item, error) {
	var items []models.Item
	result := repository.DB.Find(&items)
	return items, result.Error
}

func GetItemById(itemId uint) (models.Item, error) {
	var item models.Item
	result := repository.DB.First(&item, itemId)
	return item, result.Error
}

func InsertItem(itemData InsertItemData) error {
	result := repository.DB.Create(&models.Item{
		Name:       itemData.Name,
		ItemTypeId: itemData.ItemTypeId,
		Data:       itemData.Data,
		Price:      itemData.Price,
	})
	return result.Error
}

func UpdateItem(itemId uint, itemData InsertItemData) error {
	result := repository.DB.Model(&models.Item{}).Where("id = ?", itemId).Updates(models.Item{
		Name:       itemData.Name,
		ItemTypeId: itemData.ItemTypeId,
		Data:       itemData.Data,
		Price:      itemData.Price,
	})
	return result.Error
}

func DeleteItem(itemId uint) error {
	result := repository.DB.Delete(&models.Item{}, itemId)
	return result.Error
}
