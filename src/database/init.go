package database

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

var session *r.Session

func Connect(dUri string) (err error) {
	session, err = r.Connect(r.ConnectOpts{
		Address: dUri,
	})

	return
}

func Setup() (err error) {
	err = r.DB("aridia").TableList().Contains("problems").Do(func(x r.Term) r.Term {
		return r.Branch(x, nil, r.DB("aridia").TableCreate("problems"))
	}).Exec(session)

	if err != nil {
		return
	}

	return r.DB("aridia").TableList().Contains("users").Do(func(x r.Term) r.Term {
		return r.Branch(x, nil, r.DB("aridia").TableCreate("users"))
	}).Exec(session)
}
