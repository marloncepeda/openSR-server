package main

import (
	"log"

	// swagger docs
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/scientific-logs-api/app/config"
	_ "github.com/ctreminiom/scientific-logs-api/docs"
	"github.com/ctreminiom/scientific-logs-api/models"

	"github.com/jinzhu/gorm"

	// GORM postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	config.LoadConfigurationVariables()

	//Connect to the database
	db, err := gorm.Open("postgres", config.FormatSQLConnectionURL())

	if err != nil {
		log.Panic(err)
	}

	db.LogMode(true)

	//Migrate Schemas
	db.AutoMigrate(&models.Users{})

	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
