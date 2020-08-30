package service

import (
	"hello/dto"
	"hello/models"
	"hello/repo"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//* export  handle function
func HandleArticleGet(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := repo.DB.Where("id = ?", id).First(&article).Error; err != nil {
		ErrorByServer(c, err.Error())
		return
	}
	OK(c, article)
}

func HandleArticlesGet(c *gin.Context) {
	page, err1 := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, err2 := strconv.Atoi(c.DefaultQuery("perPage", "1"))
	if err1 != nil || err2 != nil {
		ErrorByRequest(c, "参数错误")
		return
	}
	var articles []models.Article
	var count int
	if err := PageableDB(page, perPage).Find(&articles).Count(&count).Error; err != nil {
		ErrorByServer(c, err.Error())
		return
	}
	OK(c, gin.H{"count": count, "data": articles})
}

func HandleArticlePost(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		ErrorByRequest(c, err.Error())
		return
	}
	cate, err := getArticlesCate(article.Cid)
	if err != nil {
		ErrorByRequest(c, "分类错误")
		return
	}
	article.Category = append(article.Category, cate)
	if err := repo.DB.Create(&article).Error; err != nil {
		ErrorByServer(c, err.Error())
		return
	}
	OK(c, article)
}

func HandleArticlePatch(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorByRequest(c, err.Error())
		return
	}
	id := uint(idInt)
	var newArticle dto.ArticelUpdateDto
	if err := c.ShouldBindJSON(&newArticle); err != nil {
		ErrorByRequest(c, err.Error())
		return
	}
	if err := repo.DB.Model(&models.Article{Model: gorm.Model{ID: id}}).Update(newArticle).Error; err != nil {
		ErrorByServer(c, err.Error())
		return
	}
	OK(c, "修改成功")
}

func HandleArticleDelete(c *gin.Context) {
	id := c.Param("id")
	if err := repo.DB.Where("id = ?", id).Delete(models.Article{}).Error; err != nil {
		ErrorByServer(c, "删除失败,或者记录不存在")
		return
	}
	OK(c, nil)
}

//* private handle function
func getArticlesCate(cid int) (models.Category, error) {
	var cate models.Category
	err := repo.DB.First(&cate, cid).Error
	return cate, err
}
