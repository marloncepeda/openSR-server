package main

import (
	"github.com/ctreminiom/scientific-logs-api/common/database"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/scientific-logs-api/app/users"

	// swagger docs
	_ "github.com/ctreminiom/scientific-logs-api/docs"
)

func main() {

	r := gin.New()

	db := database.Connect()
	db.AutoMigrate(&users.User{})

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	users.Routes(r)

	r.Run()
}
