package router

import (
	"go-on-store/app/controller"
	"go-on-store/app/repository"
	"go-on-store/app/service"
	"go-on-store/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	db := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	cartRepo := repository.NewCartRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	cartService := service.NewCartService(cartRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	categoryController := controller.NewCategoryController(categoryService)
	cartController := controller.NewCartController(cartService)
	transactionController := controller.NewTransactionController(transactionService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	r.GET("/products", productController.GetAllProducts)
	r.GET("/products/:id", productController.GetProductByID)
	r.GET("/products/category/:categoryID", productController.GetProductsByCategory)
	r.POST("/products", productController.AddProduct)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	r.GET("/categories", categoryController.GetAllCategories)
	r.GET("/categories/:id", categoryController.GetCategoryByID)
	r.POST("/categories", categoryController.AddCategory)
	r.PUT("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	r.GET("/cart/:userID", cartController.GetCartItems)
	r.POST("/cart", cartController.AddToCart)
	r.DELETE("/cart/:cartID", cartController.RemoveFromCart)

	r.POST("/checkout", transactionController.Checkout)

	return r
}
