package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shrhawk-entertainer/go_lang_api/common"
	"github.com/shrhawk-entertainer/go_lang_api/db"
	"github.com/shrhawk-entertainer/go_lang_api/forms"
	"github.com/shrhawk-entertainer/go_lang_api/models"
	"github.com/shrhawk-entertainer/go_lang_api/serializers"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func RetrieveUser(c *gin.Context){
	userInfo := &models.GormUser{}
	err := db.GetDB().First(&userInfo)
	if err.RecordNotFound(){
		c.JSON(http.StatusOK, gin.H{"user": "No user found"})
	}else{
		profileResponse := serializers.UserResponse{
			Email: userInfo.Email,
			Gender: userInfo.Gender,
			Name: userInfo.Name,
		}
		c.JSON(http.StatusOK, gin.H{"user": profileResponse})
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
	err := db.GetDB().Create(user).Error
	if err != nil {
		dbErrorNumber := common.GetDatabaseErrorNumber(err)
		message := "internal server error"
		if dbErrorNumber == 1062 {
			message = fmt.Sprintf("email %s already exist", UserSignUpValidationForm.Email)
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError,  gin.H{"message": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	return
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
