package main

import (
	"go-jwt/controllers"
	"go-jwt/database"
	"go-jwt/repository"
	"go-jwt/router"
	"go-jwt/services"
)

func main() {
	db, err := database.StartDB()
	if err != nil {
		return
	}
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService, db)

	productRepository := repository.NewProductRepository()
	productService := services.NewProductService(productRepository, db)
	productController := controllers.NewProductController(productService, db)

	r := router.StartApp(userController, productController)
	r.Run(":9000")
}
