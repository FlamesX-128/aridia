package database

import (
	models "github.com/FlamesX-128/aridia/src/models/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func InsertPost(problem models.Post) (err error) {
	return r.DB("aridia").Table("problems").Insert(
		problem,
		r.InsertOpts{
			Conflict: "update",
		},
	).Exec(session)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func GetPosts() (problems []models.Post, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("problems").Run(session); err != nil {
		return
	}

	if err = cursor.All(&problems); err != nil {
		return
	}

	return
}

func GetPost(id string) (problem models.Post, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("problems").Get(id).Run(session); err != nil {
		return
	}

	if err = cursor.One(&problem); err != nil {
		return
	}

	return
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func ExistsPost(id string) (exists bool, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("problems").GetAll(id).Count().Eq(1).Run(session); err != nil {
		return
	}

	if err = cursor.One(&exists); err != nil {
		return
	}

	return
}
