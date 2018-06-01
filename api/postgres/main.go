package postgres

import (
	"errors"
	"log"

	"github.com/ctreminiom/scientific-logs-api/api/postgres/models"
	"github.com/go-pg/pg"
)

// Connect ...
func Connect(user, pass, addr, database string) (*pg.DB, error) {

	options := &pg.Options{User: user, Password: pass, Addr: addr, Database: database}

	db := pg.Connect(options)

	if db == nil {
		return nil, errors.New("Failed to connect to database")
	}

	log.Printf("Connection to database successful.\n")

	models.Migrate(db)

	/*
		err := db.Close()

		if err != nil {
			log.Panic(err)
		}

		if err != nil {
			return nil, errors.New("Error while closing the connection, Reason: " + err.Error())
		}
	*/

	return db, nil
}
