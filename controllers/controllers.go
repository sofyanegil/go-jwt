package controllers

import "github.com/gin-gonic/gin"

type IUserControllers interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
}

type IProductControllers interface {
	CreateProduct(c *gin.Context)
	ReadProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	ReadAllProduct(c *gin.Context)
}

var (
	appJSON = "application/json"
)
