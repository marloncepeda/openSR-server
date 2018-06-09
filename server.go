package main

import (
	"fmt"
	"log"

	// swagger docs
	_ "github.com/ctreminiom/openSR-server/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/openSR-server/api/config"
	"github.com/ctreminiom/openSR-server/api/controllers/user"
	"github.com/ctreminiom/openSR-server/api/postgres"
	"github.com/ctreminiom/openSR-server/api/security"
)

func main() {

	err := config.Load()

	if err != nil {
		log.Panic(err)
	}

	params := config.Fetch()

	db, err := postgres.Connect(params.Username, params.Password, params.Addr, params.Database)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(security.Encrypt("asdasd"))

	r := gin.New()

	r.Use(gin.Logger())

	user.Routes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
