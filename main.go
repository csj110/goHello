package main

import (
	"hello/models"
	"hello/repo"
	"hello/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// prepare the database connect
	prePareDatabase()
	defer repo.DB.Close()
	prePareRedis()
	defer repo.RDB.Close()

	// create gin route
	r := routes.CreateRoute()
	// middleware
	// r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	// register all routeGroup
	routes.RegisterRoutes()

	// start the server
	if err := r.Run(":3000"); err != nil {
		log.Fatal("app run failed")
	}
}

func prePareDatabase() {
	// init mysql database
	err := repo.InitMySql()
	if err != nil {
		log.Fatal("no access to database", err.Error())
	}
	// migrate
	repo.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Article{})
}

func prePareRedis() {
	//redis cache
	err := repo.InitRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}
}
