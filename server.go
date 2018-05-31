package main

import (
	"github.com/ctreminiom/scientific-logs-api/app/postgresql"
	"github.com/ctreminiom/scientific-logs-api/app/users"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	// swagger docs
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/scientific-logs-api/app/config"
	_ "github.com/ctreminiom/scientific-logs-api/docs"

	// GORM postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	config.LoadConfigurationVariables()

	//Connect to the database
	db := postgresql.Connect()
	postgresql.Logger()

	//Setup gin instance
	r := gin.New()
	r.Use(gin.Logger())

	//Import routes
	users.Routes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}

func createSchemas(db *pg.DB) error {

	err := db.CreateTable(&users.People{}, &orm.CreateTableOptions{Temp: false})

	if err != nil {
		return err
	}

	return nil
}
