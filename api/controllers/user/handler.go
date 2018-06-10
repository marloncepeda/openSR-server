package user

import (
	"fmt"
	"net/http"

	"github.com/ctreminiom/openSR-server/api/postgres/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type injection struct {
	db *gorm.DB
}

func (gorm *injection) create(c *gin.Context) {

	var json registerJSON

	hasTheCorrectFormat := validateJSONBody(c.ShouldBindWith(&json, binding.JSON))

	if hasTheCorrectFormat {

		fmt.Println("PASA")

		code, message := createUser(json, gorm.db)

		c.JSON(code, message)
	}

}

func (gorm *injection) delete(c *gin.Context) {

	Identification := c.Params.ByName("id")

	user := models.User{ID: Identification}

	err := user.CheckInformation(gorm.db).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	user.Delete(gorm.db)

	c.JSON(http.StatusOK, "User deleted")

}

func (gorm *injection) update(c *gin.Context) {

	Identification := c.Params.ByName("id")

	user := models.User{ID: Identification}

	err := user.CheckInformation(gorm.db).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	user.UpdateFields(gorm.db)

	c.JSON(http.StatusOK, "User updated")

}

func (gorm *injection) password(c *gin.Context) {

	Identification := c.Params.ByName("id")
	user := models.User{ID: Identification}

	err := user.CheckInformation(gorm.db).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	user.UpdatePassword(gorm.db)

	c.JSON(http.StatusOK, "User's password updtated")
}

func (gorm *injection) getUsers(c *gin.Context) {

	users := models.User{}

	data, err := users.GetUsers(gorm.db)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(200, data)

}

func (gorm *injection) getUser(c *gin.Context) {

	Identification := c.Params.ByName("id")
	user := models.User{ID: Identification}

	data, err := user.GetUser(gorm.db)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, data)

}

// ExportRoutes ...
func ExportRoutes(gin *gin.Engine, db *gorm.DB) {

	v1 := gin.Group("/api/v1.1/module/")
	{
		api := &injection{db: db}

		v1.POST("/user", api.create)

		v1.GET("/user", api.getUsers)

		v1.GET("/user/:id", api.getUsers)

		v1.PUT("/user/:id", api.update)

		v1.DELETE("/user/:id", api.delete)

		//v1.PUT("/user/password/:id", api.password)

	}

}
