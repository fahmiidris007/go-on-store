package controller

import (
	"go-on-store/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) GetProductsByCategory(ctx *gin.Context) {
	categoryID, err := strconv.Atoi(ctx.Param("categoryID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products, err := c.service.GetProductByCategoryID(uint(categoryID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": products})
}
