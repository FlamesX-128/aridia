package database

import (
	"fmt"

	models "github.com/FlamesX-128/aridia/src/models/database"
	"github.com/FlamesX-128/aridia/src/tools"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func InsertProblem(problem models.PostProblem) (err error) {
	return r.DB("aridia").Table("problems").Insert(
		tools.ExtendPostProblem(problem),
	).Exec(session)
}

func UpdateProblem(id string, problem models.PutProblem) (err error) {
	return r.DB("aridia").Table("problems").Get(id).Update(
		tools.ExtendPutProblem(problem),
	).Exec(session)
}

func GetProblems() (problems []models.GetProblem, err error) {
	cursor, err := r.DB("aridia").Table("problems").Run(session)

	if err != nil {
		fmt.Println("An error occured while fetching the problems: ", err)

		return
	}

	if err = cursor.All(&problems); err != nil {
		fmt.Println("An error occured while decoding the problems: ", err)

		return
	}

	return
}

func GetProblem(id string) (problem models.GetProblem, err error) {
	cursor, err := r.DB("aridia").Table("problems").Get(id).Run(session)

	if err != nil {
		fmt.Println("An error occured while fetching the problem: ", err)

		return
	}

	if err = cursor.One(&problem); err != nil {
		fmt.Println("An error occured while decoding the problem: ", err)

		return
	}

	return
}
