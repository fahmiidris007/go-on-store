package controller

import (
	"go-on-store/app/model"
	"go-on-store/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	service *service.CartService
}

func NewCartController(service *service.CartService) *CartController {
	return &CartController{service: service}
}

func (c *CartController) GetCartItems(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItems, err := c.service.GetCartByUserID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cartItems})
}

func (c *CartController) AddToCart(ctx *gin.Context) {
	var cart model.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.AddCart(&cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cart})
}

func (c *CartController) UpdateCart(ctx *gin.Context) {
	cartID, err := strconv.Atoi(ctx.Param("cartID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := c.service.GetCartByID(uint(cartID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateCart(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": cart})
}

func (c *CartController) RemoveFromCart(ctx *gin.Context) {
	cartID, err := strconv.Atoi(ctx.Param("cartID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := c.service.GetCartByID(uint(cartID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.DeleteCart(cart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Cart item deleted"})
}
