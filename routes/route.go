package routes

import "github.com/gin-gonic/gin"

var r *gin.Engine

func CreateRoute()*gin.Engine{
	r= gin.Default()
	return r
}

func CreateGroup(group string)*gin.RouterGroup{
	return r.Group(group)
}