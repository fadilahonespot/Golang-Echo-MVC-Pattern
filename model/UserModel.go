package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/settings"
	"fmt"
)

type UserModel struct {
	db settings.DatabaseConfig
}

func (e *UserModel) AddUser(user *entity.User) (*entity.User, error) {
	err := e.db.GetDatabaseConfig().Save(&user).Error
	if err != nil {
		fmt.Printf("[UserModel.AddUser] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to save user")
	}
	return user, nil
}

func (e *UserModel) GetUsers()(*[]entity.User, error) {
	var users []entity.User
	err := e.db.GetDatabaseConfig().Find(&users).Error
	if err != nil {
		fmt.Printf("[userModel.GetUsers] error execute query %v \n", err)
		return nil, fmt.Errorf("failed get all users")
	}
	return &users, nil
}

func (e *UserModel) GetUserById(id int)(*entity.User, error) {
	var user = entity.User{}
	err := e.db.GetDatabaseConfig().Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Printf("[UserModel.GetUserById] error execute query %v \n", err)
		return nil, fmt.Errorf("id user is not exist")
	}
	return &user, nil
}