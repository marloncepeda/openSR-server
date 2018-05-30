package users

import (
	"net/http"

	"github.com/go-pg/pg"

	"github.com/gin-gonic/gin"
)

type orm struct {
	db *pg.DB
}

func (connection *orm) register(c *gin.Context) {

	var json registerModel
	err := c.ShouldBindJSON(&json)

	if err == nil {

		newUser := People{
			ID:            "1",
			Consecutive:   "asdsadsa",
			Name:          json.Name,
			SurName:       json.Surname,
			SecondSurName: json.SecondSurname,
			Phone:         json.Phone,
			UserName:      json.Username,
			Password:      json.Password,
		}

		err := create(newUser, connection.db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"status": "log"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Routes ...
func Routes(gin *gin.Engine, db *pg.DB) {

	v1 := gin.Group("")
	{
		env := &orm{db: db}
		v1.POST("/register", env.register)
	}
}
