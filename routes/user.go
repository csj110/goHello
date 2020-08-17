package routes

import "hello/service"

func CreateUserRoute() {
	userRoute := CreateGroup("user")
	{
		userRoute.GET("/all", service.HandleUsersGet)
		userRoute.GET("/{:id unit}",service.HandleUserGet)
	}
}
