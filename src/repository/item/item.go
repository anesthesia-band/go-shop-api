package item

import (
	"go-shop-api/src/repository"
	"go-shop-api/src/repository/models"
)

type InsertItemData struct {
	Name       string
	ItemTypeId uint
	Data       string
	Price      string
}

type UpdateItemData struct {
	Name       string
	ItemTypeId uint
	Data       string
	Price      string
	Active     bool
}

func GetAll() ([]models.Item, error) {
	var items []models.Item
	result := repository.DB.Find(&items)
	return items, result.Error
}

func GetAllActive() ([]models.Item, error) {
	var items []models.Item
	result := repository.DB.Where("active = ?", true).Find(&items)
	return items, result.Error
}

func GetById(id uint) (models.Item, error) {
	var item models.Item
	result := repository.DB.First(&item, id)
	return item, result.Error
}

func Insert(data InsertItemData) error {
	result := repository.DB.Create(&models.Item{
		Name:       data.Name,
		ItemTypeId: data.ItemTypeId,
		Data:       data.Data,
		Price:      data.Price,
	})
	return result.Error
}

func UpdateById(id uint, data UpdateItemData) error {
	result := repository.DB.Model(&models.Item{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Name":       data.Name,
		"ItemTypeId": data.ItemTypeId,
		"Data":       data.Data,
		"Price":      data.Price,
		"Active":     data.Active,
	})
	return result.Error
}

func DeleteById(id uint) error {
	result := repository.DB.Delete(&models.Item{}, id)
	return result.Error
}
