package itemtype

import (
	"go-shop-api/src/repository"
	"go-shop-api/src/repository/models"
)

type InsertItemTypeData struct {
	Title string
	Name  string
}

func GetAll() ([]models.ItemType, error) {
	var itemTypes []models.ItemType
	result := repository.DB.Find(&itemTypes)
	return itemTypes, result.Error
}

func GetById(id uint) (models.ItemType, error) {
	var itemType models.ItemType
	result := repository.DB.First(&itemType, id)
	return itemType, result.Error
}

func GetByTitle(title string) (models.ItemType, error) {
	var itemType models.ItemType
	result := repository.DB.Where("title = ?", title).First(&itemType)
	return itemType, result.Error
}

func Insert(data InsertItemTypeData) error {
	result := repository.DB.Create(&models.ItemType{
		Title: data.Title,
		Name:  data.Name,
	})
	return result.Error
}

func UpdateById(id uint, data InsertItemTypeData) error {
	result := repository.DB.Model(&models.ItemType{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Title": data.Title,
		"Name":  data.Name,
	})
	return result.Error
}

func DeleteById(id uint) error {
	result := repository.DB.Delete(&models.ItemType{}, id)
	return result.Error
}
