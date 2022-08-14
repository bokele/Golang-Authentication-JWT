package routes

import (
	controller "github.com/bokele/golang-auth-jwt/controllers"
)

func AuthRoutes(incomingRoutes Routes *gin.Engine) {

	incomingRoutes.POST("users/signup",  controller.Signup())
	incomingRoutes.POST("user/login",  controller.Login())
}
