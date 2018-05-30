package main

import (
	"log"
	"time"

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

	db2 := pg.Connect(&pg.Options{

		Addr:     config.FormatSQLAddressConnection(),
		User:     config.GetDatabaseUsername(),
		Password: config.GetDatabasePassword(),
		Database: config.GetDatabaseDatabase(),
	})

	defer db2.Close()

	db2.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	err := createSchemas(db2)

	if err != nil {
		log.Panic(err)
	}

	r := gin.New()

	users.Routes(r, db2)

	//users.Routes(r, db)
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
