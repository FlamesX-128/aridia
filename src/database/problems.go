package database

import (
	models "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/FlamesX-128/aridia/src/tools"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func InsertProblem(problem models.PostProblem) (err error) {
	return r.DB("aridia").Table("problems").Insert(
		tools.ExtendPostProblem(problem),
	).Exec(session)
}

func UpdateProblem(problem models.PutProblem) (err error) {
	return r.DB("aridia").Table("problems").Get(problem.ID).Update(
		tools.ExtendPutProblem(problem),
	).Exec(session)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func GetProblems() (problems []models.GetProblem, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("problems").Run(session); err != nil {
		return
	}

	if err = cursor.All(&problems); err != nil {
		return
	}

	return
}

func GetProblem(id string) (problem models.GetProblem, err error) {
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

func ExistsProblem(id string) (exists bool, err error) {
	var cursor *r.Cursor

	if cursor, err = r.DB("aridia").Table("problems").GetAll(id).Count().Eq(1).Run(session); err != nil {
		return
	}

	if err = cursor.One(&exists); err != nil {
		return
	}

	return
}
