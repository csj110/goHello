package service

import (
	"hello/models"
	"hello/repo"

	"github.com/gin-gonic/gin"
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
		ErrorByRequest(c, err.Error())
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
	OK(c, cate)
}

func checkCateExist(cate models.Category) bool {
	var count int
	repo.DB.Where(cate).Count(&count)
	return count > 0
}
