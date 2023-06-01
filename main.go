package main

import (

	"github.com/AD-Singh-S-Barlow/KiraKillsRecord/handlers"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()

	namesGroup := router.Group("/names")
	{
		namesGroup.GET("", handlers.GetNames)
		namesGroup.PUT("", handlers.PutNames)
		namesGroup.POST("", handlers.PostNames)
		namesGroup.DELETE("", handlers.DeleteNames)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}

	

	router.Run()
}
