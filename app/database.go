package app

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// GORM postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Open ....
func Open() {

	var db *gorm.DB
	var err error

	db, err = gorm.Open("postgres", "")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	fmt.Println(db)

}
