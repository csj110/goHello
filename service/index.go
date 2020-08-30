package service

import (
	"hello/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func PageableDB(page, perPage int) *gorm.DB {
	return repo.DB.Limit(perPage).Offset((page - 1) * perPage)
}
