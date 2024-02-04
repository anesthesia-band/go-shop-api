package grouptype

import (
	"go-shop-api/src/repository"
	"go-shop-api/src/repository/models"
)

type InsertGroupTypeData struct {
	Title string
	Name  string
}

func GetAll() ([]models.GroupType, error) {
	var groupType []models.GroupType
	result := repository.DB.Find(&groupType)
	return groupType, result.Error
}

func GetById(id uint) (models.GroupType, error) {
	var groupType models.GroupType
	result := repository.DB.First(&groupType, id)
	return groupType, result.Error
}

func GetByTitle(title string) (models.GroupType, error) {
	var groupType models.GroupType
	result := repository.DB.Where("title = ?", title).First(&groupType)
	return groupType, result.Error
}

func Insert(data InsertGroupTypeData) error {
	result := repository.DB.Create(&models.GroupType{
		Title: data.Title,
		Name:  data.Name,
	})
	return result.Error
}

func UpdateById(id uint, data InsertGroupTypeData) error {
	result := repository.DB.Model(&models.GroupType{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Title": data.Title,
		"Name":  data.Name,
	})
	return result.Error
}

func DeleteById(id uint) error {
	result := repository.DB.Delete(&models.GroupType{}, id)
	return result.Error
}
