package users

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

// CreateUser is a gin handler that creates a new user
func (c *Context) CreateUser(api *gin.Context) {

	var json template

	hasTheCorrectFormat := check(api.ShouldBindBodyWith(&json, binding.JSON))

	if hasTheCorrectFormat {

		code, message := createUser(json, c.db())

		api.JSON(code, message)
		return
	}

	api.JSON(400, "Invalid format")

}

// DeleteUser is a gin handler that creates a new user
func (c *Context) DeleteUser(api *gin.Context) {
	code, response := deleteUser(api.Params.ByName("public_id"), c.db())

	api.JSON(code, response)
}

// UpdateUser is a gin handler that creates a new user
func (c *Context) UpdateUser(api *gin.Context) {

	params := make(map[string]string, 2)

	params["id"] = api.Param("public_id")
	params["field"] = api.Param("field")
	params["value"] = api.Param("value")

	code, response := updateUser(params, c.db())

	api.JSON(code, response)

}

// GetAllUsers is a gin handler that creates a new user
func (c *Context) GetAllUsers(api *gin.Context) {

	users, err := getAllUsers(c.db())

	if err != nil {
		api.JSON(500, "Internal error")
		return
	}

	api.JSON(200, users)
}

// GetUser ...
func (c *Context) GetUser(api *gin.Context) {

	user, err := getUser(api.Param("public_id"), c.db())

	if err != nil {
		api.JSON(400, "User does not exists")
	}

	api.JSON(200, user)
}
