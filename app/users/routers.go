package users

import (
	"github.com/ctreminiom/scientific-logs-api/common/database"
	"github.com/gin-gonic/gin"
)

// SignIn ...
func SignIn(c *gin.Context) {

	var user User

	c.BindJSON(&user)

	db := database.Connect()

	db.Create(&user)

	c.JSON(200, user)

}

// Login ...
func Login(c *gin.Context) {

}

// Routes ...
func Routes(gin *gin.Engine) {

	v1 := gin.Group("api/v1")
	{
		v1.POST("/sign-in", SignIn)

		v1.GET("/login", Login)
	}

}
