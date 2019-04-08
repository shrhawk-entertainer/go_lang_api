package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/db"
	"github.com/vsouza/go-gin-boilerplate/models"
	"net/http"

	//"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func RetrieveUser(c *gin.Context){
	userInfo := db.GetDB().First(&models.GormUser{})
	c.JSON(http.StatusOK, gin.H{"user": userInfo.Value})
}
//func (u UserController) Retrieve(c *gin.Context) {
//	if c.Param("id") != "" {
//		user, err := userModel.GetByID(c.Param("id"))
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
//			c.Abort()
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
//		return
//	}
//	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
//	c.Abort()
//	return
//}
