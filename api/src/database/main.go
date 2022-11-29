package database

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type Database struct {
	Session *r.Session
	Name    string
}
