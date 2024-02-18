package controller

import (
	"go-on-store/app/model"
	"go-on-store/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": categories})
}

func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	category, err := c.service.GetCategoryByID(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": category})

}

func (c *CategoryController) AddCategory(ctx *gin.Context) {
	var category service.CategoryService
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.AddCategory(&model.Category{}); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"data": category})
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateCategory(&category); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": category})
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteCategory(uint(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": "Category deleted"})
}
