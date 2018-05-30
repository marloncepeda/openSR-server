package users

import (
	"log"
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
			ID:            2,
			Consecutive:   "asdsadsa",
			Name:          json.Name,
			SurName:       json.Surname,
			SecondSurName: json.SecondSurname,
			Phone:         json.Phone,
			UserName:      json.Username,
			Password:      json.Password,
		}

		err := connection.db.Insert(&newUser)

		if err != nil {
			log.Panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"status": newUser})
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

/*
type orm struct {
	db *gorm.DB
}

func (o *orm) signUp(c *gin.Context) {

	var json registerModel
	err := c.ShouldBindJSON(&json)

	if err == nil {

		newUser := People{
			ID:            2,
			Consecutive:   "asdsadsa",
			Name:          json.Name,
			SurName:       json.Surname,
			SecondSurName: json.SecondSurname,
			Phone:         json.Phone,
			UserName:      json.Username,
			Password:      json.Password}

		test := encrypt(newUser)

		o.db.Create(&test)

		c.JSON(http.StatusOK, gin.H{"status": test})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (o *orm) login(c *gin.Context) {

	var json loginModel
	err := c.ShouldBindJSON(&json)

	if err == nil {

	}

}

// Login ...
func Login(c *gin.Context) {

	var json loginModel
	err := c.ShouldBindJSON(&json)

	if err == nil {

		if json.User == "manu" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "logged"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

}

// Routes ...
func Routes(gin *gin.Engine, db *gorm.DB) {

	v1 := gin.Group("")
	{
		env := &orm{db: db}
		v1.POST("/signUp", env.signUp)
		v1.GET("/login", env.login)
	}
}

*/
