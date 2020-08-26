package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data, "success": true})
}

func ErrorByRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{"msg": err})
}
func ErrorByServer(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
}
