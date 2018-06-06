package main

import (
	"fmt"
	"log"

	// swagger docs
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/openSR-server/api/config"
	"github.com/ctreminiom/openSR-server/api/controllers/user"
	"github.com/ctreminiom/openSR-server/api/postgres"

	_ "github.com/ctreminiom/openSR-server/docs"
)

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
