package main

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// swagger docs
	_ "github.com/ctreminiom/scientific-logs-api/docs"

	"github.com/gin-gonic/gin"

	// GORM postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	r := gin.New()

	// Load configuration files
	//yaml, err := app.Config()

	//if err != nil {
	//	log.Fatal(err)
	//}

	// Connect to the database
	//db, _ := gorm.Open("postgres", app.FormatSQLConnectionURL(yaml))

	//if err != nil {
	//	log.Panic(err)
	//}

	//Init API

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

	/*

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

	*/
}
