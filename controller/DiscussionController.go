package controller

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/model"
	"Golang-Echo-MVC-Pattern/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type DiscussionController struct {
	discussionModel  model.DiscussionModel
	userModel        model.UserModel
	catagoryModel    model.CatagoryModel
	uploadController UploadController
}

func (e *DiscussionController) AddDiscussion(c echo.Context) error {
	var discussion = entity.Discussion{}
	err := c.Bind(&discussion)
	if err != nil {
		fmt.Printf("[DiscussionController.AddDiscussion] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	_, err = e.catagoryModel.GetCatagoryId(int(discussion.CatagoryID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	_, err = e.userModel.GetUserById(int(discussion.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	if discussion.Message == "" || discussion.Title == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	mDiscussion, err := e.discussionModel.AddDiscussion(&discussion)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, mDiscussion)
}

func (e *DiscussionController) AddDiscussionFirst(c echo.Context) error {
	var discussionFirst = entity.DiscussionFirst{}
	err := c.Bind(&discussionFirst)
	if err != nil {
		fmt.Printf("[DiscussionController.AddDiscussionFirst] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oppss server someting wrong")
	}
	_, err = e.discussionModel.GetDiscussionById(int(discussionFirst.DiscussionID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	_, err = e.userModel.GetUserById(int(discussionFirst.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	if discussionFirst.Message == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	output, err := e.discussionModel.AddDiscussionFirst(&discussionFirst)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, output)
}

func (e *DiscussionController) GetAllDiscussion(c echo.Context) error {
	discussions, err := e.discussionModel.GetDiscussions()
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	var total int
	var discussionShow []entity.DiscussionShow
	for i := 0; i < len(*discussions); i++ {
		user, err := e.userModel.GetUserById(int((*discussions)[i].UserID))
		if err != nil {
			return utils.HandleError(c, http.StatusNotFound, err.Error())
		}
		discussionFirst, err := e.discussionModel.GetDiscussionFirstByDiscussionId(int((*discussions)[i].ID))
		if err != nil {
			return utils.HandleError(c, http.StatusNotFound, err.Error())
		}
		total += len(*discussionFirst)
		for k := 0; k < len(*discussionFirst); k++ {
			discussionSecond, err := e.discussionModel.GetDiscussionSecondByDiscussionFirstId(int((*discussionFirst)[k].ID))
			if err != nil {
				return utils.HandleError(c, http.StatusNotFound, err.Error())
			}
			total += len(*discussionSecond)
		}
		var res = entity.DiscussionShow{
			ID:      (*discussions)[i].ID,
			Name:    user.Name,
			Date:    (*discussions)[i].CreatedAt.String(),
			Message: (*discussions)[i].Message,
			Total:   total,
		}
		discussionShow = append(discussionShow, res)
		total = 0
	}
	return utils.HandleSuccess(c, discussionShow)
}

func (e *DiscussionController) GetDiscussionSecondByDiscussionFirstID(id int) (*[]entity.DiscussionSecondDetailShow, error) {
	discussionSecond, err := e.discussionModel.GetDiscussionSecondByDiscussionFirstId(id)
	if err != nil {
		return nil, err
	}
	var discussionSecondShow []entity.DiscussionSecondDetailShow
	for i := 0; i < len(*discussionSecond); i++ {
		user, err := e.userModel.GetUserById(int((*discussionSecond)[i].UserID))
		if err != nil {
			return nil, err
		}
		var res = entity.DiscussionSecondDetailShow{
			ID:      (*discussionSecond)[i].ID,
			Name:    user.Name,
			Date:    (*discussionSecond)[i].CreatedAt.String(),
			Message: (*discussionSecond)[i].Message,
		}
		discussionSecondShow = append(discussionSecondShow, res)
	}
	return &discussionSecondShow, nil
}

func (e *DiscussionController) GetDiscussionFirstByDiscussionId(id int) (*[]entity.DiscussionFirstDetailShow, error) {
	discussionFirst, err := e.discussionModel.GetDiscussionFirstByDiscussionId(id)
	if err != nil {
		return nil, err
	}
	var discussionFirstShow []entity.DiscussionFirstDetailShow
	for i := 0; i < len(*discussionFirst); i++ {
		user, err := e.userModel.GetUserById(int((*discussionFirst)[i].UserID))
		if err != nil {
			return nil, err
		}
		discussionSecond, err := e.GetDiscussionSecondByDiscussionFirstID(int((*discussionFirst)[i].ID))
		if err != nil {
			return nil, err
		}
		var res = entity.DiscussionFirstDetailShow{
			ID:               (*discussionFirst)[i].ID,
			Name:             user.Name,
			Date:             (*discussionFirst)[i].CreatedAt.String(),
			Message:          (*discussionFirst)[i].Message,
			SecondDiscussion: (*discussionSecond),
		}
		discussionFirstShow = append(discussionFirstShow, res)
	}
	return &discussionFirstShow, nil
}

func (e *DiscussionController) GetDiscussionDetailById(c echo.Context) error {
	idStr := c.Param("discussionId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, "id has be number")
	}
	discussion, err := e.discussionModel.GetDiscussionById(id)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	user, err := e.userModel.GetUserById(int(discussion.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	discussionFirstShow, err := e.GetDiscussionFirstByDiscussionId(id)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	files, err := e.discussionModel.GetDiscussionFilesByDiscussionId(id)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	var partFiles []string
	for i := 0; i < len(*files); i++ {
		partFiles = append(partFiles, (*files)[i].File)
	}
	images, err := e.discussionModel.GetDiscussionImagesByDiscussionId(id)
	var partImages []string
	for k := 0; k < len(*images); k++ {
		partImages = append(partImages, (*images)[k].Image)
	}
	var discussionShow = entity.DiscussionDetailShow{
		ID:             discussion.ID,
		Name:           user.Name,
		Date:           discussion.CreatedAt.String(),
		Message:        discussion.Message,
		FirsDiscussion: (*discussionFirstShow),
		Files:          partFiles,
		Images:         partImages,
	}
	return utils.HandleSuccess(c, discussionShow)
}

func (e *DiscussionController) AddDiscussionSecond(c echo.Context) error {
	var discussionSecond = entity.DiscussionSecond{}
	err := c.Bind(&discussionSecond)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, "Oopsss server someting wrong")
	}

	if discussionSecond.UserID == 0 || discussionSecond.DiscussionFirstID == 0 || discussionSecond.Message == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	_, err = e.userModel.GetUserById(int(discussionSecond.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	_, err = e.discussionModel.GetDiscussionFirstById(int(discussionSecond.DiscussionFirstID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	output, err := e.discussionModel.AddDiscussionSecond(&discussionSecond)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, output)
}

func (e *DiscussionController) AddFilesImagesDiscussion(c echo.Context) error {
	discussionIdSrt := c.FormValue("discussion_id")
	discussionId, err := strconv.Atoi(discussionIdSrt)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, "id has be number")
	}
	userIdStr := c.FormValue("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, "user id has be number")
	}

	tx := e.discussionModel.BeginTrans()

	discussion, err := e.discussionModel.GetDiscussionById(discussionId)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	user, err := e.userModel.GetUserById(userId)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	partFiles, err := e.uploadController.UploadFiles(c)
	if err != nil {
		return utils.HandleError(c, http.StatusForbidden, err.Error())
	}
	for i := 0; i < len(partFiles); i++ {
		var res = entity.Files{
			UserID:       uint(userId),
			DiscussionID: uint(discussionId),
			File:         partFiles[i],
		}
		_, err := e.discussionModel.AddFileDiscussion(&res, tx)
		if err != nil {
			utils.RollbackFiles(partFiles)
			tx.Rollback()
			return utils.HandleError(c, http.StatusInternalServerError, err.Error())
		}
	}
	partImages, err := e.uploadController.UploadImages(c)
	if err != nil {
		utils.RollbackFiles(partFiles)
		return utils.HandleError(c, http.StatusForbidden, err.Error())
	}
	for i := 0; i < len(partImages); i++ {
		var res = entity.Images{
			UserID:       uint(userId),
			DiscussionID: uint(discussionId),
			Image:        partImages[i],
		}
		_, err := e.discussionModel.AddImageDiscussion(&res, tx)
		if err != nil {
			utils.RollbackFiles(partFiles)
			utils.RollbackFiles(partImages)
			tx.Rollback()
			return utils.HandleError(c, http.StatusInternalServerError, err.Error())
		}
	}
	var discussionPost = entity.DiscussionPost{
		ID:      discussion.ID,
		UserID:  user.ID,
		Title:   discussion.Title,
		Date:    discussion.CreatedAt.String(),
		Message: discussion.Message,
		Images:  partImages,
		Files:   partFiles,
	}
	tx.Commit()
	return utils.HandleSuccess(c, discussionPost)
}

func (e *DiscussionController) EditDiscussion(c echo.Context) error {
	idStr := c.Param("discussionId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, "id discussion has be number")
	}
	var discussion = entity.Discussion{}
	_, err = e.discussionModel.GetDiscussionById(id)
	err = c.Bind(&discussion)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, "Oopss server someting wrong")
	}
	_, err = e.catagoryModel.GetCatagoryId(int(discussion.CatagoryID))
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err.Error())
	}
	upDiscussion, err := e.discussionModel.UpdateDiscussion(id, &discussion)
	if err != nil {
		return utils.HandleError(c, http.StatusForbidden, "Opss server someting wrong")
	}
	return utils.HandleSuccess(c, upDiscussion)
}

func (e *DiscussionController) DeleteDiscussionById(c echo.Context) error {
	idStr := c.Param("discussionId")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, "id has be number")
	}
	tx := e.discussionModel.BeginTrans()
	discussionFiles, err := e.discussionModel.GetDiscussionFilesByDiscussionId(id)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	var partFiles []string
	for i := 0; i < len(*discussionFiles); i++ {
		partFiles = append(partFiles, (*discussionFiles)[i].File)
	}
	discussionImages, err := e.discussionModel.GetDiscussionImagesByDiscussionId(id)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	var partImages []string
	for k := 0; k < len(*discussionImages); k++ {
		partImages = append(partImages, (*discussionImages)[k].Image)
	}
	err = e.discussionModel.DeleteDiscussionFilesByDiscussionId(id, tx)
	if err != nil {
		tx.Rollback()
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	err = e.discussionModel.DeleteDiscussionImagesByDiscussionId(id, tx)
	if err != nil {
		tx.Rollback()
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	discussionFirst, err := e.discussionModel.GetDiscussionFirstByDiscussionId(id)
	if err != nil {
		tx.Rollback()
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	for i := 0; i < len(*discussionFirst); i++ {
		err := e.discussionModel.DeleteDiscussionSecondByDiscussionFirstId(int((*discussionFirst)[i].ID), tx)
		if err != nil {
			tx.Rollback()
			return utils.HandleError(c, http.StatusNotFound, err.Error())
		}
	}
	err = e.discussionModel.DeleteDiscussionFirstByDiscussionId(id, tx)
	if err != nil {
		tx.Rollback()
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	err = e.discussionModel.DeleteDicussionById(id, tx)
	if err != nil {
		tx.Rollback()
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	utils.RollbackFiles(partFiles)
	utils.RollbackFiles(partImages)
	tx.Commit()
	return utils.HandleSuccess(c, "Successs delete discussion")
}
