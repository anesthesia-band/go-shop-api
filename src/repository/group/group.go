package group

import (
	"go-shop-api/src/repository"
	"go-shop-api/src/repository/models"
)

type InsertGroupData struct {
	Name        string
	GroupTypeId uint
	Data        string
	Price       string
}

type UpdateGroupData struct {
	Name        string
	GroupTypeId uint
	Data        string
	Price       string
	Active      bool
}

func GetAll() ([]models.Group, error) {
	var groups []models.Group
	result := repository.DB.Find(&groups)
	return groups, result.Error
}

func GetAllActive() ([]models.Group, error) {
	var groups []models.Group
	result := repository.DB.Where("active = ?", true).Find(&groups)
	return groups, result.Error
}

func GetById(id uint) (models.Group, error) {
	var groups models.Group
	result := repository.DB.First(&groups, id)
	return groups, result.Error
}

func Insert(data InsertGroupData) error {
	result := repository.DB.Create(&models.Group{
		Name:        data.Name,
		GroupTypeId: data.GroupTypeId,
		Data:        data.Data,
		Price:       data.Price,
	})
	return result.Error
}

func UpdateById(id uint, data UpdateGroupData) error {
	result := repository.DB.Model(&models.Group{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Name":        data.Name,
		"GroupTypeId": data.GroupTypeId,
		"Data":        data.Data,
		"Price":       data.Price,
		"Active":      data.Active,
	})
	return result.Error
}

func DeleteById(id uint) error {
	result := repository.DB.Delete(&models.Group{}, id)
	return result.Error
}
