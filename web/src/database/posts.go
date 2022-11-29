package database

import (
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func InsertPost(post mdb.Post) (mdb.Post, error) {
	rows, err := r.DB("aridia").Table("posts").Insert(
		post, r.InsertOpts{
			Conflict:      "update",
			ReturnChanges: true,
		},
	).Run(session)

	if err != nil {
		return post, err
	}

	var nvalue map[string]interface{}

	if err = rows.One(&nvalue); err != nil {
		return post, err
	}

	post.Id = nvalue["generated_keys"].([]interface{})[0].(string)

	return post, nil
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func GetPosts(filter map[string]interface{}) (posts []mdb.Post, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("posts").Filter(filter).Run(session); err != nil {
		return
	}

	if err = cursor.All(&posts); err != nil {
		return
	}

	return
}

func GetPost(id string) (post mdb.Post, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("posts").Get(id).Run(session); err != nil {
		return
	}

	if err = cursor.One(&post); err != nil {
		return
	}

	return
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func ExistsPost(id string) (exists bool, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("posts").GetAll(id).Count().Eq(1).Run(session); err != nil {
		return
	}

	if err = cursor.One(&exists); err != nil {
		return
	}

	return
}
