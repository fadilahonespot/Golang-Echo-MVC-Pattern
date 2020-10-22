package entity

import "github.com/jinzhu/gorm"

type DiscussionFirst struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"user_id"`
	DiscussionID uint   `gorm:"not null" json:"discussion_id"`
	Message      string `gorm:"type:text; not null" json:"message"`
}

func (e DiscussionFirst) TableName() string {
	return "discussion_first"
}
