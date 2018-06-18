package database

import (
	"log"

	"github.com/ctreminiom/openSR-server/pkg/services/consecutives/types"
	"github.com/ctreminiom/openSR-server/pkg/services/users"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// Postgres gorm driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Init creates a connection to postgres database and migrate the models
func Init() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", viper.Get("Connection"))

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	err = migrate(db)

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {

	err := db.AutoMigrate(&users.Model{}, &types.Model{}).Error

	if err != nil {
		return err
	}

	return nil
}
