package routes

import (
	"github.com/ctreminiom/openSR-server/pkg/auth"
	"github.com/ctreminiom/openSR-server/pkg/services/consecutives/types"
	"github.com/ctreminiom/openSR-server/pkg/services/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Export return the gin pointer that has all api routes
func Export(c *gin.Engine, db *gorm.DB) {

	user := c.Group("/api/v1")
	user.Use(auth.AuthorizationHandler())
	{
		gin := users.Context{Pool: db}

		user.POST("/users", gin.CreateUser)

		user.GET("/users", gin.GetAllUsers)

		user.GET("/users/:public_id", gin.GetUser)

		user.DELETE("/users/:public_id", gin.DeleteUser)

		user.PUT("/users/:public_id/:field/:value", gin.UpdateUser)
	}

	consecutiveType := c.Group("/api/v1")
	consecutiveType.Use(auth.AuthorizationHandler())
	{
		gin := types.Context{Pool: db}

		consecutiveType.POST("/consecutive/types", gin.CreateType)
		consecutiveType.GET("/consecutive/types", gin.GetTypes)
		consecutiveType.GET("/consecutive/types/:id", gin.GetType)
	}

	login := c.Group("/api/v1")
	{
		gin := users.Context{Pool: db}

		login.GET("/login", gin.Login)
	}

}
