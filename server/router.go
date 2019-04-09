package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shrhawk-entertainer/go_lang_api/app_environments"
	"github.com/shrhawk-entertainer/go_lang_api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.BasicAuth(app_environments.BASIC_AUTH_CREDENTAILS))

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userGroup.GET("/:id", controllers.RetrieveUser)
			userGroup.POST("/", controllers.CreateUser)
			userGroup.DELETE('/', controllers.DeleteUser)
		}
	}
	return router

}
