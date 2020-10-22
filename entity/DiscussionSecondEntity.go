package entity

import "github.com/jinzhu/gorm"

type DiscussionSecond struct {
	gorm.Model
	UserID            uint   `gorm:"not null" json:"user_id"`
	DiscussionFirstID uint   `gorm:"not null" json:"discussion_first_id"`
	Message           string `gorm:"type:text; not null" json:"message"`
}

func (e DiscussionSecond) TableName() string {
	return "discussion_second"
}
