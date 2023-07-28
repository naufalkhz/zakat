package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
)

type UserInterface interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
}

type userImplementation struct {
	svc services.UserService
}

func NewUserInterface(svc services.UserService) UserInterface {
	return &userImplementation{
		svc: svc,
	}
}

func (e *userImplementation) Get(c *gin.Context) {

	user, err := e.svc.Get(c)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(c, http.StatusOK, user)
}

func (e *userImplementation) Create(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := e.svc.Create(c, user)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(c, http.StatusOK, user)
}
