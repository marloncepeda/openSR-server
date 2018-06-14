package main

import (
	"github.com/ctreminiom/openSR-server/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := api.Start()

	// Add middlewares
	r.Use(gin.Logger())

	//Execute it
	r.Run()

}
