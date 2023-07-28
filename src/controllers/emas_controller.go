package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/services"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, resultEmas)
}
