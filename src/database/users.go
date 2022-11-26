package database

import (
	models "github.com/FlamesX-128/aridia/src/models/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func InsertUser(id string) (err error) {
	return r.DB("aridia").Table("users").Insert(models.User{
		ID:        id,
		Permisson: 0,
	}).Exec(session)
}

func GetUser(id string) (user models.User, err error) {
	cursor, err := r.DB("aridia").Table("users").Get(id).Run(session)

	if err != nil {
		return
	}

	if err = cursor.One(&user); err != nil {
		return
	}

	return
}
