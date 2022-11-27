package database

import (
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/FlamesX-128/aridia/src/tools"
	"golang.org/x/oauth2"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func CreateUser(id string, auth oauth2.Token) error {
	return r.DB("aridia").Table("users").Insert(
		tools.CreateUser(id, auth),
	).Exec(session)
}

func UpdateUserPost(uId string, pId string) error {
	return r.DB("aridia").Table("users").Filter(map[string]string{
		"Id": uId,
	}).Update(map[string]interface{}{
		"Posts": r.Row.Field("Posts").Append(pId),
	}).Exec(session)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func GetUsers() (users []mdb.SimplifiedUser, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("users").Run(session); err != nil {
		return
	}

	if err = cursor.All(&users); err != nil {
		return
	}

	return
}

func GetUser(id string) (user mdb.User, err error) {
	var cursor *r.Cursor

	// the id is not a uuid, so we need to get the user by the id field
	if cursor, err = r.DB("aridia").Table("users").Filter(map[string]string{
		"Id": id,
	}).Run(session); err != nil {
		return
	}

	if err = cursor.One(&user); err != nil {
		return
	}

	return
}
