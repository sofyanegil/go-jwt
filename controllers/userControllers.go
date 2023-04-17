package controllers

import (
	"go-jwt/helpers"
	"go-jwt/models"
	"go-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService services.IUserService
	DB          *gorm.DB
}

func NewUserController(userService services.IUserService, db *gorm.DB) *UserController {
	return &UserController{
		UserService: userService,
		DB:          db,
	}
}

func (controller UserController) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	userReturn, err := controller.UserService.Register(&User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        userReturn.ID,
		"email":     userReturn.Email,
		"full_name": userReturn.FullName,
		"role":      userReturn.Role,
	})
}
func (controller UserController) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	token, err := controller.UserService.Login(User.Email, User.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invaild email/password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
