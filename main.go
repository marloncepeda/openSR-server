package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/scientific-logs-api/app/users"

	// swagger docs
	_ "github.com/ctreminiom/scientific-logs-api/docs"

	// GORM postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	r := gin.New()

	db, err := gorm.Open("postgres", "host=ec2-18-188-250-148.us-east-2.compute.amazonaws.com port=5432 user=ctreminio dbname=science sslmode=disable password=ctreminio")

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	// Export Models
	users.Export(db)

	// Routes
	users.Routes(r, db)

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
