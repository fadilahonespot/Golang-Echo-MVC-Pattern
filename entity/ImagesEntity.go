package entity

import "github.com/jinzhu/gorm"

type Images struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	DiscussionID uint   `gorm:"not null" json:"discusion_id"`
	Image        string `gorm:"not null" json:"image"`
}

func (e Images) TableName() string {
	return "images"
}