package postgresql

import (
	"log"
	"time"

	"github.com/ctreminiom/scientific-logs-api/app/config"
	"github.com/go-pg/pg"
)

var values = config.Database()

// Connect ...
func Connect() (db *pg.DB) {

	db = pg.Connect(&pg.Options{
		Addr:     values["addr"],
		User:     values["username"],
		Password: values["password"],
		Database: values["database"],
	})

	return db
}

// Logger ....
func Logger() {

	db := Connect()

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {

		query, err := event.FormattedQuery()

		if err != nil {
			log.Panic(err)
		}

		log.Printf("%s %s", time.Since(event.StartTime), query)

	})

}
