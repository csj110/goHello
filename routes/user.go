package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateUserRoute() {
	userRoute := CreateGroup("user")
	{
		userRoute.GET("/", func(c *gin.Context) {

		})
	}
}
