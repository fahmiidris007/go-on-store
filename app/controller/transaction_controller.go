package controller

import (
	"go-on-store/app/model"
	"go-on-store/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *service.TransactionService
}

func NewTransactionController(service *service.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) Checkout(ctx *gin.Context) {
	var transaction model.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.AddTransaction(&transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": transaction})
}
