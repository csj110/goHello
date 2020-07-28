package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/models"
	"hello/repo"
	"hello/util"
	"net/http"
	"strings"
)

func AuthGuard(genUser bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, "Bearer ")
		if !(len(parts) == 2) {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userId", mc.UserId)
		if genUser {
			var user models.User
			if err:=repo.DB.Find(&user).Error;err!=nil{
				c.JSON(http.StatusOK, gin.H{
					"code": 2005,
					"msg":  "无效的Token",
				})
				fmt.Println("jwt secret 可能已经泄露")
				c.Abort()
				return
			}
			c.Set("user",user)
		}
		c.Next() // 后续的处理函数可以用过c.Get("username")
	}
}
