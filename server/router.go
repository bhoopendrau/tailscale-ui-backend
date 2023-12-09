package server

import (
	"github.com/bhoopendrau/tailscale-ui-backend/controllers"
	"github.com/bhoopendrau/tailscale-ui-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.POST("/", user.SignUp)
		}
	}
	return router

}
