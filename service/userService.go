package service

import (
	"github.com/gin-gonic/gin"
	"hello/models"
	"hello/repo"
	"net/http"
	"strconv"
)

// get pageable users
func HandleUsersGet(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	perPage, err := strconv.Atoi(c.Query("perPage"))
	if err != nil {
		page = 1
		perPage = 10
	}
	if page < 0 {
		page = 1
	}
	if perPage < 0 {
		perPage = 10
	}
	var users []models.User
	if err = repo.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

//get single user
func HandleUserGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err = repo.DB.Where("id = ?", id).Delete(models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
