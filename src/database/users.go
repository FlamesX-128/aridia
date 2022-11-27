package database

import (
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func CreateUser(user mdb.User) (err error) {
	return r.DB("aridia").Table("users").Insert(
		user,
	).Exec(session)
}

func UpdateUser(user mdb.User) (err error) {
	return r.DB("aridia").Table("users").Get(user.ID).Update(
		user,
	).Exec(session)
}

func GetUser(id string) (user mdb.User, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("users").Get(id).Run(session); err != nil {
		return
	}

	if err = cursor.One(&user); err != nil {
		return
	}

	return
}
