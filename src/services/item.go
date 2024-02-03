package services

import (
	"go-shop-api/src/models"
)

type InsertItemData struct {
	Name       string
	ItemTypeId uint `gorm:"index"`
	Data       string
	Price      string
}

func GetItems() ([]models.Item, error) {
	var items []models.Item
	result := db.Find(&items)
	return items, result.Error
}

func GetItemById(itemId uint) (models.Item, error) {
	var item models.Item
	result := db.First(&item, itemId)
	return item, result.Error
}

func InsertItem(itemData InsertItemData) error {
	result := db.Create(&models.Item{
		Name:       itemData.Name,
		ItemTypeId: itemData.ItemTypeId,
		Data:       itemData.Data,
		Price:      itemData.Price,
	})
	return result.Error
}

func UpdateItem(itemId uint, itemData InsertItemData) error {
	result := db.Model(&models.Item{}).Where("id = ?", itemId).Updates(models.Item{
		Name:       itemData.Name,
		ItemTypeId: itemData.ItemTypeId,
		Data:       itemData.Data,
		Price:      itemData.Price,
	})
	return result.Error
}

func DeleteItem(itemId uint) error {
	result := db.Delete(&models.Item{}, itemId)
	return result.Error
}
