package main

import (
	"fmt"

	// swagger docs

	"github.com/ctreminiom/scientific-logs-api/app"
	_ "github.com/ctreminiom/scientific-logs-api/docs"

	// GORM postgres driver

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//r := gin.New()

	// Load configuration files
	yaml, _ := app.Config()

	fmt.Println(yaml)

	//fmt.Println(security.EncryptWithAES("Caros Treminio"))

	// Connect to the database
	//db, _ := gorm.Open("postgres", app.FormatSQLConnectionURL(yaml))

	//fmt.Println(db)

	//if err != nil {
	//	log.Panic(err)
	//}

	//Init API

	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//	r.Run()

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
