package controllers

import (
	"go-jwt/helpers"
	"go-jwt/models"
	"go-jwt/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService services.IProductService
	DB             *gorm.DB
}

func NewProductController(productService services.IProductService, db *gorm.DB) *ProductController {
	return &ProductController{
		ProductService: productService,
		DB:             db,
	}
}

func (controller ProductController) CreateProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	productReturn, err := controller.ProductService.CreateProduct(&Product, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, productReturn)
}

func (controller ProductController) ReadAllProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	Product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	productReturn, err := controller.ProductService.ReadAllProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productReturn)
}

func (controller ProductController) ReadProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	productReturn, err := controller.ProductService.ReadProduct(productId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productReturn)
}

func (controller ProductController) UpdateProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	productReturn, err := controller.ProductService.UpdateProduct(&Product, productId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productReturn)
}

func (controller ProductController) DeleteProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := controller.ProductService.DeleteProduct(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
