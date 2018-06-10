package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// Postgres gorm driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect ...
func Connect() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", viper.Get("Connection"))

	if err != nil {
		return nil, err
	}

	return db, nil
}
