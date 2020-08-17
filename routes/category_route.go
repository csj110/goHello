package routes

import "hello/service"

func CreateCategoryRoute()  {
		CategoryRoute:=CreateGroup("cate")
		{
			CategoryRoute.GET("/all",service.HandleCategoriesGet)
			CategoryRoute.POST("",service.HandleCategoryPost)
		}
}