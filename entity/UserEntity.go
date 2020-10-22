package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name    string `gorm:"not null" json:"name"`
	Email   string `gorm:"not null" json:"email"`
}

func (e User) TableName() string {
	return "user"
}