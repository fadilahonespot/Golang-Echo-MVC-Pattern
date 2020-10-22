package entity

import "github.com/jinzhu/gorm"

type Discussion struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id"`
	CatagoryID uint   `gorm:"not null" json:"catagory_id"`
	Title      string `gorm:"not null" json:"title"`
	Message    string `gorm:"type:text; not null" json:"message"`
}

func (e Discussion) TableName() string {
	return "discussion"
}
