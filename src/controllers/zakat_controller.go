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
	CreateZakatTabungan(c *gin.Context)
	CreateZakatPerdagangan(c *gin.Context)
	CreateZakatEmas(c *gin.Context)
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

	utils.SuccessResponse(c, http.StatusOK, zakatPenghasilanRes)
}

func (e *zakatImplementation) CreateZakatTabungan(c *gin.Context) {
	var zakatTabungan *models.ZakatTabunganRequest
	if err := c.ShouldBindJSON(&zakatTabungan); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	zakatTabunganRes, err := e.svc.CreateTabungan(c, zakatTabungan)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, zakatTabunganRes)
}
func (e *zakatImplementation) CreateZakatPerdagangan(c *gin.Context) {
	var zakatPerdagangan *models.ZakatPerdaganganRequest
	if err := c.ShouldBindJSON(&zakatPerdagangan); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	zakatPerdaganganRes, err := e.svc.CreatePerdagangan(c, zakatPerdagangan)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, zakatPerdaganganRes)
}

func (e *zakatImplementation) CreateZakatEmas(c *gin.Context) {
	var zakatEmas *models.ZakatEmasRequest
	if err := c.ShouldBindJSON(&zakatEmas); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	zakatEmasRes, err := e.svc.CreateEmas(c, zakatEmas)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, zakatEmasRes)
}
