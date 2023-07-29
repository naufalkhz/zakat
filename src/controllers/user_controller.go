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
	Edit(c *gin.Context)
	GetUserSessionRest(c *gin.Context)
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
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, user)
}

func (e *userImplementation) Create(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := e.svc.Create(c, user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, user)
}

func (e *userImplementation) Edit(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := e.svc.Edit(c, user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, user)
}

func (e *userImplementation) GetUserSession(c *gin.Context) (*models.User, error) {
	user, err := e.svc.GetUserSession(c)
	if err != nil {
		return nil, err
	}
	return user, err
}

// TODO: Remove this function, just for checking on rest
func (e *userImplementation) GetUserSessionRest(c *gin.Context) {
	user, err := e.svc.GetUserSession(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, nil)
	}
	utils.SuccessResponse(c, http.StatusOK, user)
}
