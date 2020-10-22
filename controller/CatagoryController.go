package controller

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/model"
	"Golang-Echo-MVC-Pattern/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type CatagoryController struct {
	catagoryModel model.CatagoryModel
	userModel     model.UserModel
}

func (e *CatagoryController) AddCatagory(c echo.Context) error {
	var catagory = entity.Catagory{}
	err := c.Bind(&catagory)
	if err != nil {
		fmt.Printf("[CatagoryController.AddCatagory] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oopss server has be wrong")
	}

	_, err = e.userModel.GetUserById(int(catagory.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	mCatagory, err := e.catagoryModel.AddCatagory(&catagory)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, mCatagory)
}

func (e *CatagoryController) ViewAllCatagory(c echo.Context) error {
	catagories, err := e.catagoryModel.GetCatagories()
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, catagories)
}
