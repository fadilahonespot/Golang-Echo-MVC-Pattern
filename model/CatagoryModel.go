package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/settings"
	"fmt"
)


type CatagoryModel struct {
	db settings.DatabaseConfig
}

func (e *CatagoryModel) AddCatagory(catagory *entity.Catagory) (*entity.Catagory, error) {
	err := e.db.GetDatabaseConfig().Save(&catagory).Error
	if err != nil {
		fmt.Printf("[CatagortModel.AddCatagory] error execute query %v \n", err)
		return nil, fmt.Errorf("Failed add data catagory")
	}
	return catagory, nil
}

func (e *CatagoryModel) GetCatagories() (*[]entity.Catagory, error) {
	var catagories []entity.Catagory
	err := e.db.GetDatabaseConfig().Find(&catagories).Error
	if err != nil {
		fmt.Printf("[catagoryModel.GetCatagories] error execute query %v \n", err)
		return nil, fmt.Errorf("Failed show all catagory")
	}
	return &catagories, nil
}

func (e *CatagoryModel) GetCatagoryId(id int)(*entity.Catagory, error) {
	var catagory = entity.Catagory{}
	err := e.db.GetDatabaseConfig().Table("catagory").Where("id = ?", id).First(&catagory).Error
	if err != nil {
		fmt.Printf("[CatagoryModel.getCatagoryId] error execute query %v \n", err)
		return nil, fmt.Errorf("Id catagory is not exist")
	}
	return &catagory, nil
}