package router

import (
	"go-jwt/controllers"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp(userControllers controllers.IUserControllers, productController controllers.IProductControllers) *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", userControllers.UserRegister)
		userRouter.POST("/login", userControllers.UserLogin)

	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/", productController.ReadAllProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), productController.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), productController.ReadProduct)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), productController.DeleteProduct)
	}

	return r
}
