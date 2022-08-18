package controller

import (
	"gin-boiler-plate/model"
	"gin-boiler-plate/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	FindById(*gin.Context)
	FindAll(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController() UserController {
	return &userController{userService: service.NewUserService()}
}

func (u *userController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.userService.FindById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *userController) FindAll(c *gin.Context) {
	users, err := u.userService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *userController) Create(c *gin.Context) {
	createUser := new(model.User)
	err := c.BindJSON(&createUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	createdUser, err := u.userService.Create(*createUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, createdUser)
}

func (u *userController) Update(c *gin.Context) {
	updateUser := new(model.User)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = c.BindJSON(&updateUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	updatedUser, err := u.userService.Update(uint(id), *updateUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, updatedUser)
}

type ErrorMessage struct {
	error
}

func (u *userController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = u.userService.Delete(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}
