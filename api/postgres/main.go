package postgres

import (
	"errors"
	"log"

	"github.com/ctreminiom/openSR-server/api/postgres/models"
	"github.com/go-pg/pg"
)

// Connect ...
func Connect(user, pass, addr, database string) (*pg.DB, error) {

	options := &pg.Options{User: user, Password: pass, Addr: addr, Database: database}

	db := pg.Connect(options)

	if db == nil {
		return nil, errors.New("Failed to connect to database")
	}

	err := models.Migrate(db)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Connection to database successful.\n")

	return db, nil
}
