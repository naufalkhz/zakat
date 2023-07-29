package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
)

type ZakatInterface interface {
	CreateZakatPenghasilan(c *gin.Context)
}

type zakatImplementation struct {
	svc services.ZakatService
}

func NewZakatInterface(svc services.ZakatService) ZakatInterface {
	return &zakatImplementation{
		svc: svc,
	}
}

func (e *zakatImplementation) CreateZakatPenghasilan(c *gin.Context) {
	var zakatPenghasilan *models.ZakatPenghasilanRequest
	if err := c.ShouldBindJSON(&zakatPenghasilan); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	zakatPenghasilanRes, err := e.svc.CreatePenghasilan(c, zakatPenghasilan)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, zakatPenghasilanRes)
}
