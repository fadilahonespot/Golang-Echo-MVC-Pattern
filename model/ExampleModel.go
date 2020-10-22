package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/settings"
)

type ExampleModel struct {
	db settings.DatabaseConfig
}

// Get Example Post
func (ExampleModel ExampleModel) GetPosts() []entity.ExampleEntity {
	posts := []entity.ExampleEntity{
		{
			Title:       "NewsOne",
			Description: "NewsOneDescription",
		},
		{
			Title:       "NewsTwo",
			Description: "NewsTwoDescription",
		},
	}

	return posts
}


