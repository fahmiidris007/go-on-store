package router

import (
	"go-on-store/app/controller"
	"go-on-store/app/middleware"
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

	protected := r.Group("/")
	protected.Use(middleware.Middleware())
	{
		protected.GET("/products", productController.GetAllProducts)
		protected.GET("/products/:id", productController.GetProductByID)
		protected.GET("/products/category/:categoryID", productController.GetProductsByCategory)
		protected.POST("/products", productController.AddProduct)
		protected.PUT("/products/:id", productController.UpdateProduct)
		protected.DELETE("/products/:id", productController.DeleteProduct)

		protected.GET("/categories", categoryController.GetAllCategories)
		protected.GET("/categories/:id", categoryController.GetCategoryByID)
		protected.POST("/categories", categoryController.AddCategory)
		protected.PUT("/categories/:id", categoryController.UpdateCategory)
		protected.DELETE("/categories/:id", categoryController.DeleteCategory)

		protected.GET("/cart/:userID", cartController.GetCartItems)
		protected.POST("/cart", cartController.AddToCart)
		protected.PUT("/cart/:cartID", cartController.UpdateCart)
		protected.DELETE("/cart/:cartID", cartController.RemoveFromCart)

		protected.GET("/checkout/:userID", transactionController.GetTransactionByUserID)
		protected.POST("/checkout", transactionController.Checkout)
	}
	return r
}
