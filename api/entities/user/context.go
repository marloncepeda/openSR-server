package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type context struct {
	pool *gorm.DB
}

func (c *context) db() *gorm.DB { return c.pool }

//CRUD callbacks
func (c *context) create(api *gin.Context) {

	var json creation

	hasTheCorrectFormat := check(api.ShouldBindWith(&json, binding.JSON))

	if hasTheCorrectFormat {

		code, response := create(json, c.db())

		api.JSON(code, response)

	}
	api.JSON(400, "Invalid format")

}

func (c *context) delete(api *gin.Context) {
	code, response := erase(api.Params.ByName("id"), c.db())

	api.JSON(code, response)
}

func (c *context) update(api *gin.Context) {
	code, response := update(api.Params.ByName("id"), c.db())

	api.JSON(code, response)
}

func (c *context) getUsers(api *gin.Context) {

	code, users, response := getUsers(c.db())

	if users == nil {
		api.JSON(code, response)
	}

	api.JSON(code, users)
}

func (c *context) getUser(api *gin.Context) {

	code, user, response := getUser(c.db())

	if user.ID == "error" {
		api.JSON(code, response)
	}

	api.JSON(code, user)

}
