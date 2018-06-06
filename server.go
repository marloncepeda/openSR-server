package main

import (
	"fmt"
	"log"

	// swagger docs
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/scientific-logs-api/api/config"
	"github.com/ctreminiom/scientific-logs-api/api/controllers/user"
	"github.com/ctreminiom/scientific-logs-api/api/postgres"

	_ "github.com/ctreminiom/scientific-logs-api/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {

	err := config.Load()

	if err != nil {
		fmt.Println(err.Error())
	}

	params := config.Fetch()

	db, err := postgres.Connect(params.Username, params.Password, params.Addr, params.Database)

	if err != nil {
		log.Panic(err)
	}

	r := gin.New()
	r.Use(gin.Logger())

	user.Routes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
