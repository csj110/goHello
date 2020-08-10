package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/models"
	"hello/repo"
	"hello/util"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func HandleGetInfo(c *gin.Context) {
	if user, ok := c.Get("user"); ok {

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{})
}

func HandlePostLogin(c *gin.Context) {
	var loginDto models.LoginDto
	if err := c.ShouldBind(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "errBinding",
		})
		return
	}
	var user models.User
	if err := repo.DB.Debug().FirstOrCreate(&user, map[string]interface{}{"phone": loginDto.Phone}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	tokenString, err := util.GenToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func HandlePostCaptcha(c *gin.Context)  {
	var captchaDto models.CaptchaDto
	if err:=c.ShouldBind(&captchaDto);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}
	var code = strconv.Itoa(rand.Intn(10000)+10000)[1:]
	fmt.Println(code)
	err:=repo.SetExp(captchaDto.Phone+":au",code,time.Minute*5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{"success":true})
}