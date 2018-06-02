package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ctreminiom/scientific-logs-api/api/postgres/models"
	"github.com/go-pg/pg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type orm struct {
	database *pg.DB
}

func (connection orm) register(c *gin.Context) {

	var json registerTemplate

	isCompleted := validate(c.ShouldBindWith(&json, binding.JSON))

	if isCompleted {

		// Count All user in the database
		count, _ := connection.database.Model((*models.User)(nil)).Count()

		fmt.Println(strconv.Itoa(count + 1))

		fmt.Println(encrypt(strconv.Itoa(count + 1)))

		user := models.User{
			ID:            encrypt(strconv.Itoa(count + 1)),
			Consecutive:   encrypt("i"),
			Name:          encrypt(json.Name),
			Surname:       encrypt(json.Surname),
			SecondSurName: encrypt(json.SecondSurname),
			Phone:         encrypt(json.Phone),
			Username:      encrypt(json.Username),
			Password:      encrypt(json.Password),
		}

		err := user.Save(connection.database)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user created succefully"})

	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameters"})

}

// Routes ...
func Routes(gin *gin.Engine, db *pg.DB) {

	v1 := gin.Group("")
	{
		env := &orm{database: db}
		v1.POST("/register", env.register)
	}
}
