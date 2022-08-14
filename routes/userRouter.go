package routes

import (
	controller "github.com/bokele/golang-auth-jwt/controllers"
)

func UserhRoutes(incomingRoutes Routes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("users", controller.GetUsers())
	incomingRoutes.POST("user/login", controller.GetUser())
}
