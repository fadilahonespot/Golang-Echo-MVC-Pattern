package entity

import "github.com/jinzhu/gorm"

type Catagory struct {
	gorm.Model
	UserID   uint   `gorm:"not null" json:"user_id"`
	Catagory string `gorm:"not null" json:"catagory"`
}

func (e Catagory) TableName() string {
	return "catagory"
}
