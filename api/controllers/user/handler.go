package user

import (
	"net/http"

	"github.com/ctreminiom/scientific-logs-api/api/security/jwt"
	"github.com/go-pg/pg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type orm struct {
	db *pg.DB
}

func (connection orm) register(c *gin.Context) {

	var json registerTemplate

	isValidated := validateHTTPBody(c.ShouldBindWith(&json, binding.JSON))

	if isValidated {

		code, value := create(json, connection.db)

		c.JSON(code, gin.H{"message": value})
		return

	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameters"})
	return

}

func (connection orm) login(c *gin.Context) {

	var json loginTemplate

	isValidated := validateHTTPBody(c.ShouldBindWith(&json, binding.JSON))

	if isValidated {

		bolean, _ := confirmUsername(json.Username, json.Password, connection.db)

		if bolean {
			c.JSON(http.StatusOK, gin.H{"message": jwt.Encode(json.Username)})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"message": "usernae and password incorrect"})
		return

	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameters"})
	return

}

// Routes ...
func Routes(gin *gin.Engine, db *pg.DB) {

	v1 := gin.Group("")
	{
		env := &orm{db: db}
		v1.POST("/register", env.register)
		v1.POST("/login", env.login)

	}
}
