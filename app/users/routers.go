package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type orm struct {
	db *gorm.DB
}

func (o *orm) signUp(c *gin.Context) {

	var json registerModel
	err := c.ShouldBindJSON(&json)

	if err == nil {

		newUser := Users{
			Consecutive:   12,
			Name:          json.Name,
			SurName:       json.Surname,
			SecondSurName: json.SecondSurname,
			Phone:         json.Phone,
			UserName:      json.Username,
			Password:      json.Password}

		o.db.Create(&newUser)

		c.JSON(http.StatusOK, gin.H{"status": newUser})

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
