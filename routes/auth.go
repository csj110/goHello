package routes

import (
	"hello/middleware"
	"hello/service"
)

func CreateAuthRoute() {
	authRoute := CreateGroup("auth")
	{
		authRoute.POST("login", service.HandlePostLogin)
		authRoute.GET("info", middleware.AuthGuard(true), service.HandleGetInfo)
	}
}
