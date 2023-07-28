package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
)

type AuthInterface interface {
	SignIn(c *gin.Context)
}

type authImplementation struct {
	svc services.AuthService
}

func NewAuthInterface(svc services.AuthService) AuthInterface {
	return &authImplementation{
		svc: svc,
	}
}

func (e *authImplementation) SignIn(c *gin.Context) {
	var auth *models.AuthRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		utils.SendResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := e.svc.Login(c, auth)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	tokenString, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendResponse(c, http.StatusOK, models.AuthReponse{Token: tokenString, User: *user})
}
