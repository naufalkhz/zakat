package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
	"github.com/spf13/cast"
)

type UserInterface interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	// GetUserSession(c *gin.Context) (*models.User, error)
	GetUserSession(c *gin.Context)
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

func UserIdSession(context *gin.Context) (uint, error) {
	tokenString := context.GetHeader("Authorization")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed parse token")
	}

	var user = make(map[string]interface{})

	for key, val := range claims {
		user[key] = val
	}

	return cast.ToUint(user["id"]), nil
}

func (e *userImplementation) GetUserSession(c *gin.Context) {
	idUser, err := UserIdSession(c)

	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, nil)
	}

	// var user *models.User
	user, err := e.svc.GetUserById(c, idUser)
	if err != nil {
		utils.SendResponse(c, http.StatusInternalServerError, nil)
	}
	utils.SendResponse(c, http.StatusOK, user)
}
