package main

import (

	// swagger docs
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/ctreminiom/scientific-logs-api/docs"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

}
