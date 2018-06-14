package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Routes export the endpoint user has.
func Routes(gin *gin.Engine, db *gorm.DB) {

	CRUD := gin.Group("/api/v1/module/")
	{
		c := &context{pool: db}

		CRUD.GET("/auth", c.login)
	}
}
