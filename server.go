package main

import (
	"log"

	"github.com/ctreminiom/openSR-server/api/controllers/user"
	"github.com/ctreminiom/openSR-server/api/postgres/models"

	// swagger docs
	_ "github.com/ctreminiom/openSR-server/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/openSR-server/api/config"
	"github.com/ctreminiom/openSR-server/api/postgres"
)

func main() {

	err := config.LoadEnvironmentVariables()

	if err != nil {
		log.Panic(err)
	}

	db, err := postgres.Connect()

	if err != nil {
		log.Panic(err)
	}

	err = models.Migrate(db)
	if err != nil {
		log.Panic(err)
	}

	r := gin.New()

	r.Use(gin.Logger())

	user.ExportRoutes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
