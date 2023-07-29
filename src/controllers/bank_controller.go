package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
	"github.com/spf13/cast"
)

type BankInterface interface {
	GetListBank(c *gin.Context)
	GetBankById(c *gin.Context)
	Create(c *gin.Context)
}

type bankImplementation struct {
	svc services.BankService
}

func NewBankInterface(svc services.BankService) BankInterface {
	return &bankImplementation{
		svc: svc,
	}
}

func (e *bankImplementation) GetListBank(c *gin.Context) {
	res, err := e.svc.GetListBank(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, res)
}

func (e *bankImplementation) Create(c *gin.Context) {
	var bank *models.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := e.svc.Create(c, bank)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, res)
}

func (e *bankImplementation) GetBankById(c *gin.Context) {
	id := cast.ToUint(c.Param("id"))

	if id == 0 {
		utils.ErrorResponse(c, http.StatusBadRequest, "id parameter must be numeric")
		return
	}

	res, err := e.svc.GetBankById(c, id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, res)
}
