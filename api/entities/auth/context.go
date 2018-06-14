package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type context struct {
	pool *gorm.DB
}

func (c *context) db() *gorm.DB { return c.pool }

func (c *context) login(api *gin.Context) {

	code, response := validate(api.GetHeader("Authorization"), c.db())
	api.JSON(code, response)
}
