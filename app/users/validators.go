package users

import (
	"github.com/go-pg/pg"
)

func create(user People, db *pg.DB) error {

	userEncrypted := encrypt(user)

	err := db.Insert(userEncrypted)

	if err != nil {
		return err
	}

	return nil

}
