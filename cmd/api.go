package main

import (
	"log"

	"github.com/ctreminiom/openSR-server/configs"
	"github.com/ctreminiom/openSR-server/pkg/database"
	"github.com/ctreminiom/openSR-server/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration variables
	err := configs.Load()

	if err != nil {
		log.Panic(err)
	}

	//Database
	db, err := database.Init()

	if err != nil {
		log.Panic(err)
	}

	//Load Gin API

	api := gin.New()

	// Add middlewares
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	//Export routes
	routes.Export(api, db)

	//Execute it
	api.Run()

}
