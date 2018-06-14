package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	// swagger docs
	_ "github.com/ctreminiom/openSR-server/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/ctreminiom/openSR-server/api/config"
	"github.com/ctreminiom/openSR-server/api/entities/auth"
	"github.com/ctreminiom/openSR-server/api/entities/user"
	"github.com/ctreminiom/openSR-server/api/postgres"
	"github.com/ctreminiom/openSR-server/api/postgres/models"
)

// Load configuration variables
func load() {
	err := config.Init()

	if err != nil {
		log.Panic(err)
	}
}

// Open the PostgreSQL connection pools
func connect() *gorm.DB {

	db, err := postgres.Connect()

	if err != nil {
		log.Panic(err)
	}

	return db
}

// Export the API routes
func export(gin *gin.Engine, db *gorm.DB) {

	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.Routes(gin, db)

	auth.Routes(gin, db)

}

// Migrate the database tables
func migrate(db *gorm.DB) {

	err := models.Migrate(db)

	if err != nil {
		log.Panic(err)
	}

}

// Start function load the API configuration and launch the API
func Start() *gin.Engine {

	load()

	db := connect()

	migrate(db)

	engine := gin.New()

	export(engine, db)

	return engine
}
