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

type UserController struct {
	userModel model.UserModel
}

func (e *UserController) AddUserController(c echo.Context) error {
	var user = entity.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("[UserController.AddUser] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, "Oppss server someting wrong")
	}
	if user.Email == "" || user.Name == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	mUser, err := e.userModel.AddUser(&user)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, mUser)
}

func (e *UserController) ViewUsersController(c echo.Context) error {
	users, err := e.userModel.GetUsers()
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, users)
}

func (e *UserController) ViewUserIdController(c echo.Context) error {
	idStr := c.Param("idUser")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, "id has be number")
	}
	user, err := e.userModel.GetUserById(id)
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	return utils.HandleSuccess(c, user)
}
