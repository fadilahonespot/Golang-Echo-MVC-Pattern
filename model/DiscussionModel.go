package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/settings"
	"fmt"

	"github.com/jinzhu/gorm"
)

type DiscussionModel struct {
	db settings.DatabaseConfig
}

func (e *DiscussionModel) BeginTrans() *gorm.DB {
	return e.db.GetDatabaseConfig().Begin()
}

func (e *DiscussionModel) AddDiscussion(discussion *entity.Discussion) (*entity.Discussion, error) {
	err := e.db.GetDatabaseConfig().Save(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.AddDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data discussion")
	}
	return discussion, nil
}

func (e *DiscussionModel) GetDiscussions() (*[]entity.Discussion, error) {
	var discussions []entity.Discussion
	err := e.db.GetDatabaseConfig().Find(&discussions).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetDiscussions] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data discussion")
	}
	return &discussions, nil
}

func (e *DiscussionModel) GetDiscussionById(id int) (*entity.Discussion, error) {
	var discussion = entity.Discussion{}
	err := e.db.GetDatabaseConfig().Table("discussion").Where("id = ?", id).First(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetdiscussionById] error execute query %v \n", err)
		return nil, fmt.Errorf("id discussion is not exist")
	}
	return &discussion, nil
}

func (e *DiscussionModel) AddDiscussionFirst(discussionFirst *entity.DiscussionFirst) (*entity.DiscussionFirst, error) {
	err := e.db.GetDatabaseConfig().Save(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.AddDiscussionFirs] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to add data discussion")
	}
	return discussionFirst, nil
}

func (e *DiscussionModel) GetDiscussionFirstById(id int) (*entity.DiscussionFirst, error) {
	var discussionFirst = entity.DiscussionFirst{}
	err := e.db.GetDatabaseConfig().Table("discussion_first").Where("id = ?", id).First(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetDiscussionFirstById] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion first id is not exist")
	}
	return &discussionFirst, nil
}

func (e *DiscussionModel) GetDiscussionFirstByDiscussionId(id int) (*[]entity.DiscussionFirst, error) {
	var discussionFirst []entity.DiscussionFirst
	err := e.db.GetDatabaseConfig().Table("discussion_first").Where("discussion_id = ?", id).Find(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[discussionModel.getDiscussionFirstBydiscussionId] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion id is not exist")
	}
	return &discussionFirst, nil
}

func (e *DiscussionModel) GetDiscussionFirstAll() (*[]entity.DiscussionFirst, error) {
	var discussionFirst []entity.DiscussionFirst
	err := e.db.GetDatabaseConfig().Find(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetDiscussionFirstAll] error execute query %v \n", err)
		return nil, fmt.Errorf("Failed to show all data discussion first")
	}
	return &discussionFirst, nil
}

func (e *DiscussionModel) AddDiscussionSecond(discussionSecond *entity.DiscussionSecond) (*entity.DiscussionSecond, error) {
	err := e.db.GetDatabaseConfig().Save(&discussionSecond).Error
	if err != nil {
		fmt.Printf("[discussionModel.AddDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data discussion second")
	}
	return discussionSecond, nil
}

func (e *DiscussionModel) GetDiscussionSecondByDiscussionFirstId(id int) (*[]entity.DiscussionSecond, error) {
	var discussionSecond []entity.DiscussionSecond
	err := e.db.GetDatabaseConfig().Table("discussion_second").Where("discussion_first_id = ?", id).Find(&discussionSecond).Error
	if err != nil {
		fmt.Printf("[discussionModel.GetdiscussionSecondByDiscussionFirstId] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion first id is not exist")
	}
	return &discussionSecond, nil
}

func (e *DiscussionModel) AddFileDiscussion(file *entity.Files, tx *gorm.DB) (*entity.Files, error) {
	err := tx.Save(&file).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.AddFileDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to add file to discussion")
	}
	return file, nil
}

func (e *DiscussionModel) AddImageDiscussion(image *entity.Images, tx *gorm.DB) (*entity.Images, error) {
	err := tx.Save(&image).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.AddImageDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to add image to discussion")
	}
	return image, nil
}

func (e *DiscussionModel) GetDiscussionFilesByDiscussionId(id int) (*[]entity.Files, error) {
	var files []entity.Files
	err := e.db.GetDatabaseConfig().Table("files").Where("discussion_id = ?", id).Find(&files).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetFilesByDiscussionId] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion id is not exist")
	}
	return &files, nil
}

func (e *DiscussionModel) GetDiscussionImagesByDiscussionId(id int) (*[]entity.Images, error) {
	var images []entity.Images
	err := e.db.GetDatabaseConfig().Table("images").Where("discussion_id = ?", id).Find(&images).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetImagesByDiscussionId] error execute query %v \n", err)
		return nil, fmt.Errorf("discussion id is not exist")
	}
	return &images, nil
}

func (e *DiscussionModel) UpdateDiscussion(id int, discussion *entity.Discussion) (*entity.Discussion, error) {
	var upDiscussion = entity.Discussion{}
	err := e.db.GetDatabaseConfig().Table("discussion").Where("id = ?", id).First(&upDiscussion).Update(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.UpdateDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed to update discussion")
	}
	return &upDiscussion, nil
}

func (e *DiscussionModel) DeleteDicussionById(id int, tx *gorm.DB) error {
	var discussion = entity.Discussion{}
	err := e.db.GetDatabaseConfig().Table("discussion").Where("id = ?", id).First(&discussion).Delete(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.DeleteDiscussion] error execute query %v \n", err)
		return fmt.Errorf("discussion id is not exist")
	}
	return nil
}

func (e *DiscussionModel) DeleteDiscussionFirstByDiscussionId(id int, tx *gorm.DB) error {
	var discussionFirst []entity.DiscussionFirst
	err := tx.Table("discussion_first").Where("discussion_id = ?", id).Find(&discussionFirst).Delete(&discussionFirst).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.DeleteDiscussionFirstByDiscussionId] error execute query %v \n", err)
		return fmt.Errorf("discussion id is not exist")
	}
	return nil
}

func (e *DiscussionModel) DeleteDiscussionSecondByDiscussionFirstId(id int, tx *gorm.DB) error {
	var discussionSecond []entity.Discussion
	err := tx.Table("discussion_second").Where("discussion_first_id = ?", id).Find(&discussionSecond).Delete(&discussionSecond).Error
	if err != nil {
		fmt.Printf("[discussionModel.DeleteDiscussionSecondByDiscussionFirstId] error execute query %v \n", err)
		return fmt.Errorf("discussion first id is not exist")
	}
	return nil
}

func (e *DiscussionModel) DeleteDiscussionFilesByDiscussionId(id int, tx *gorm.DB) error {
	var discussionFiles []entity.Files
	err := tx.Table("files").Where("discussion_id = ?", id).Find(&discussionFiles).Delete(&discussionFiles).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.DeleteDiscussionFilesByDiscussionId] error execute query %v \n", err)
		return fmt.Errorf("discussion id is not exist")
	}
	return nil
}

func (e *DiscussionModel) DeleteDiscussionImagesByDiscussionId(id int, tx *gorm.DB) error {
	var discussionImages []entity.Images
	err := tx.Table("images").Where("discussion_id = ?", id).Find(&discussionImages).Delete(&discussionImages).Error
	if err != nil {
		fmt.Printf("[discussionModel.DeleteDiscussionImagesByDiscussionId] error execute query %v \n", err)
		return fmt.Errorf("discussion id is not exist")
	}
	return nil
}
