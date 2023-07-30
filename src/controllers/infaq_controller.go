package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
	"github.com/spf13/cast"
)

type InfaqInterface interface {
	CreateInfaq(c *gin.Context)
	GetListInfaq(c *gin.Context)
	CreateInfaqRiwayat(c *gin.Context)
	GetListInfaqRiwayatLastLimit(c *gin.Context)
}

type infaqImplementation struct {
	svc services.InfaqService
}

func NewInfaqInterface(svc services.InfaqService) InfaqInterface {
	return &infaqImplementation{
		svc: svc,
	}
}

func (e *infaqImplementation) CreateInfaq(c *gin.Context) {
	var infaq *models.Infaq
	if err := c.ShouldBindJSON(&infaq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	infaqRes, err := e.svc.CreateInfaq(c, infaq)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, infaqRes)
}

func (e *infaqImplementation) GetListInfaq(c *gin.Context) {
	res, err := e.svc.GetList(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, res)
}

func (e *infaqImplementation) CreateInfaqRiwayat(c *gin.Context) {
	var infaqRequest *models.InfaqRiwayatRequest
	if err := c.ShouldBindJSON(&infaqRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	infaqRes, err := e.svc.CreateInfaqRiwayat(c, infaqRequest)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, infaqRes)
}

func (e *infaqImplementation) GetListInfaqRiwayatLastLimit(c *gin.Context) {
	limit := cast.ToInt(c.Param("limit"))
	if limit > 30 {
		limit = 30
	}
	res, err := e.svc.GetRiwayatInfaqLastLimit(c, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, res)
}
