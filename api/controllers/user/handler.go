package user

import (
	"fmt"
	"net/http"

	"github.com/go-pg/pg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type orm struct {
	db *pg.DB
}

func (connection orm) register(c *gin.Context) {

	var json registerTemplate

	isValidated := validate(c.ShouldBindWith(&json, binding.JSON))

	if isValidated {

		code, value := create(json, connection.db)

		c.JSON(code, gin.H{"message": value})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameters"})

	}

}

func (connection orm) login(c *gin.Context) {

	var json loginTemplate

	isValidated := validate(c.ShouldBindWith(&json, binding.JSON))

	if isValidated {

		username := decrypt(json.Username)

		fmt.Println(username)

	}

}

// Routes ...
func Routes(gin *gin.Engine, db *pg.DB) {

	v1 := gin.Group("")
	{
		env := &orm{db: db}
		v1.POST("/register", env.register)
		v1.POST("/login", env.register)

	}
}
