package main

import (
	"hello/models"
	"hello/repo"
	"hello/routes"
	"log"
)

func main() {
	// prepare the database connect
	prePareDatabase()
	defer repo.DB.Close()

	// create gin route
	r := routes.CreateRoute()
	// register routeGroup
	routes.CreateAuthRoute()
	routes.CreateUserRoute()

	// serve the server
	if err := r.Run(":3000"); err != nil {
		log.Fatal("app run failed")
	}
}

func prePareDatabase(){
	// init mysql database
	err := repo.InitMySql()
	if err != nil {
		log.Fatal("no access to database",err.Error())
	}
	// migrate
	repo.DB.AutoMigrate(&models.User{})
}