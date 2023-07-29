package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
)

type EmasInterface interface {
	Get(c *gin.Context)
}

type emasImplementation struct {
	svc services.EmasService
}

func NewEmasInterface(svc services.EmasService) EmasInterface {
	return &emasImplementation{
		svc: svc,
	}
}

func (e *emasImplementation) Get(c *gin.Context) {
	resultEmas, err := e.svc.Get(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, resultEmas)
}
