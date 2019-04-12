package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shrhawk-entertainer/go_lang_api/common"
	"github.com/shrhawk-entertainer/go_lang_api/db"
	"github.com/shrhawk-entertainer/go_lang_api/forms"
	"github.com/shrhawk-entertainer/go_lang_api/models"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func RetrieveUser(c *gin.Context){
	userInfo := db.GetDB().First(&models.GormUser{})
	if userInfo.RecordNotFound(){
		c.JSON(http.StatusOK, gin.H{"user": "No user found"})
	}else{
		c.JSON(http.StatusOK, gin.H{"user": userInfo.Value})
	}
	return
}

func CreateUser(c *gin.Context){
	var UserSignUpValidationForm = forms.UserSignupValidation{}
	if err := c.ShouldBind(&UserSignUpValidationForm); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}
	gender := UserSignUpValidationForm.Gender
	if gender == "" {
		gender = "female"
	}
	user := &models.GormUser{
		Name: UserSignUpValidationForm.Name,
		Email: UserSignUpValidationForm.Email,
		Active: true,
		Gender: gender,
	}
	err := db.GetDB().Create(user)
	if err != nil {
		fmt.Println(err.Error)
		c.AbortWithStatusJSON(500,  gin.H{"message": err.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func DeleteUser(c *gin.Context){
	error_deletion := db.GetDB().Where([]int64{2}).Delete(&models.GormUser{}).Error
	if error_deletion !=nil {

	}
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
