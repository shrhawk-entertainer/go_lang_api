package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/controllers"
	"github.com/vsouza/go-gin-boilerplate/app_environments"
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
		}
	}
	return router

}
