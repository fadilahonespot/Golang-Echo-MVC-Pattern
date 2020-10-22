package entity

import "github.com/jinzhu/gorm"

type Files struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	DiscussionID uint   `gorm:"not null" json:"discussion_id"`
	File         string `gorm:"not null" json:"file"`
}

func (e Files) TableName() string {
	return "Files"
}
