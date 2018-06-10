package models

import (
	"github.com/jinzhu/gorm"
)

// Migrate ....
func Migrate(db *gorm.DB) error {

	if db.HasTable(&User{}) {
		db.DropTableIfExists(&User{})
	}

	err := db.AutoMigrate(&User{}).Error

	if err != nil {
		return err
	}

	return nil
}
