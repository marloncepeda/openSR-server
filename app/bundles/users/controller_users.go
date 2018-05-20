package users

import (
	"github.com/ctreminiom/scientific-logs-api/app"
	"github.com/gin-gonic/gin"
)

// CreateUser ....
func CreateUser(c *gin.Context) {

	var user User
	c.BindJSON(&user)

	db := app.Connect()

	db.Create(&user)
	c.JSON(200, user)
}
