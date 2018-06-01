package main

import (
	"fmt"
	"log"

	// swagger docs
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/scientific-logs-api/api/config"
	"github.com/ctreminiom/scientific-logs-api/api/postgres"
	"github.com/ctreminiom/scientific-logs-api/api/postgres/models"
	_ "github.com/ctreminiom/scientific-logs-api/docs"
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

	fmt.Println(db)

	newUser := models.User{
		ID:            "IDDDD",
		Consecutive:   "asd",
		Name:          "asd",
		Surname:       "asd",
		SecondSurName: "asd",
		Phone:         "asd",
		Username:      "asd",
		Password:      "asd",
	}

	err = newUser.Save(db)

	if err != nil {
		log.Panic(err)
	}

	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
