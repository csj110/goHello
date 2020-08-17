package routes

import "hello/service"

func CreateArticleRoute(){
	articleRoute:=CreateGroup("article")
	{
		articleRoute.GET("",service.HandleArticleGet)
		articleRoute.POST("",service.HandleArticlePost)
		articleRoute.PATCH("/:id",service.HandleArticlePatch)
		articleRoute.DELETE("/:id",service.HandleArticleDelete)
	}
}