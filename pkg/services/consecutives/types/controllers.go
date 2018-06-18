package types

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

// Context is a struct that contains tha gorm pool pointer
type Context struct {
	Pool *gorm.DB
}

func (c *Context) db() *gorm.DB { return c.Pool }

// CreateType is a gin handler that creates a new user
func (c *Context) CreateType(api *gin.Context) {

	var json template

	hasTheCorrectFormat := check(api.ShouldBindBodyWith(&json, binding.JSON))

	if hasTheCorrectFormat {

		code, message := createType(json, c.db())

		api.JSON(code, message)
		return
	}

	api.JSON(400, "Invalid format")

}

// GetTypes is a gin handler that creates a new user
func (c *Context) GetTypes(api *gin.Context) {

	users, err := getTypes(c.db())

	if err != nil {
		api.JSON(500, "Internal error")
		return
	}

	api.JSON(200, users)
}

// GetType ...
func (c *Context) GetType(api *gin.Context) {

	user, err := getType(api.Param("id"), c.db())

	if err != nil {
		api.JSON(400, "Type does not exists")
		return
	}

	api.JSON(200, user)
}
