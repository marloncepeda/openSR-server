package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Routes export the endpoint user has.
func Routes(gin *gin.Engine, db *gorm.DB) {

	CRUD := gin.Group("/api/v1/module/")
	{
		c := &context{pool: db}

		CRUD.POST("/user", c.create)

		CRUD.DELETE("/user/:id", c.delete)

		CRUD.PUT("/user/:id", c.update)

		CRUD.GET("/users", c.getUsers)

		CRUD.GET("/user/:id", c.getUser)
	}
}
