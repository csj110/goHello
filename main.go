package main

import (
	"github.com/gin-gonic/gin"
	"hello/middleware"
	"hello/models"
	"hello/repo"
	"hello/util"
	"log"
	"net/http"
)

func main() {
	// init mysql database
	err := repo.InitMySql()
	if err != nil {
		log.Fatal("no access to database",err.Error())
	}
	defer repo.DB.Close()
	// migrate
	repo.DB.AutoMigrate(&models.User{})

	// create gin route
	r := gin.Default()

	authRoute := r.Group("auth")
	{
		authRoute.POST("login", func(c *gin.Context) {
			var loginDto models.LoginDto
			if err := c.ShouldBind(&loginDto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err":"errBinding",
				})
				return
			}
			var user models.User
			if err:=repo.DB.Debug().FirstOrCreate(&user,map[string]interface{}{"phone": loginDto.Phone}).Error;err!=nil{
				c.JSON(http.StatusInternalServerError, gin.H{
					"err":err.Error(),
				})
				return
			}

			tokenString, err := util.GenToken(user.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err":err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"token": tokenString,
			})
		})

		authRoute.GET("info", middleware.AuthGuard(true), func(c *gin.Context) {
			if user, ok := c.Get("user"); ok {
				c.JSON(http.StatusOK, gin.H{
					"user": user.(*models.User),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{})

		})
	}
	userRoute := r.Group("user")
	{
		userRoute.GET("/", func(c *gin.Context) {

		})
	}

	if err := r.Run(":3000"); err != nil {
		log.Fatal("app run failed")
	}
}
