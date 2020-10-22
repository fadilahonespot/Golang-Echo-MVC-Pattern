package settings

import (
	"Golang-Echo-MVC-Pattern/entity"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DatabaseConfig struct{}


// MySql Db Config
func (DatabaseConfig DatabaseConfig) GetDatabaseConfig() *gorm.DB {
	DB, err := gorm.Open("mysql", viper.GetString("database.mysql"))
	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		entity.User{},
		entity.Discussion{},
		entity.Catagory{},
		entity.DiscussionFirst{},
		entity.DiscussionSecond{},
		entity.Images{},
		entity.Files{},
	)

	DB.Model(&entity.Discussion{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Discussion{}).AddForeignKey("catagory_id", "catagory(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Catagory{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.DiscussionFirst{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.DiscussionFirst{}).AddForeignKey("discussion_id", "discussion(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.DiscussionSecond{}).AddForeignKey("discussion_first_id", "discussion_first(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.DiscussionSecond{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Images{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Images{}).AddForeignKey("discussion_id", "discussion(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Files{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Files{}).AddForeignKey("discussion_id", "discussion(id)", "CASCADE", "CASCADE")

	return DB
}
