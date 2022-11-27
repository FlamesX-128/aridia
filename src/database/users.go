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
	return r.DB("aridia").Table("users").Get(uId).Update(
		map[string]interface{}{
			"problems": r.Row.Field("problems").Append(pId),
		},
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
