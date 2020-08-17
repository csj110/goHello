package service

import (
	"github.com/gin-gonic/gin"
	"hello/models"
	"hello/repo"
	"net/http"
)

func HandleCategoriesGet(c *gin.Context) {
	var cates []models.Category
	if err := repo.DB.Find(&cates).Error; err != nil {
		ErrorByRequest(c, err.Error())
		return
	}
	OK(c, cates)
}
func HandleCategoryPost(c *gin.Context) {
	var cate models.Category
	if err := c.ShouldBind(&cate); err != nil {
		ErrorByRequest(c,err.Error())
		return
	}
	if checkCateExist(cate) {
		ErrorByRequest(c, "this cate is already exists")
		return
	}
	if err := repo.DB.Create(&cate).Error; err != nil {
		ErrorByServer(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cate})
}

func checkCateExist(cate models.Category) bool {
	var count int
	repo.DB.Where(cate).Count(&count)
	return count > 0
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data, "success": true})
}

func ErrorByRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{"msg": err})
}
func ErrorByServer(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
}
