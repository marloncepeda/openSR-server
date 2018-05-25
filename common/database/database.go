package database

import (
	"log"

	"github.com/jinzhu/gorm"

	// GORM postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect return a *gorm pointer
func Connect() *gorm.DB {

	db, err := gorm.Open("postgres", "host=ec2-18-188-250-148.us-east-2.compute.amazonaws.com port=5432 user=ctreminio dbname=science sslmode=disable password=ctreminio")

	if err != nil {
		log.Panic("err")
	}

	defer db.Close()

	return db
}
