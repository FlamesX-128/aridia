package database

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (db *Database) UpdateElement(table string, id string, element interface{}) error {
	return r.DB(db.Name).Table(table).Get(id).Update(element).Exec(db.Session)
}

func (db *Database) CreateElement(table string, element interface{}) error {
	return r.DB(db.Name).Table(table).Insert(element).Exec(db.Session)
}

func (db *Database) DeleteElement(table string, id string) error {
	return r.DB(db.Name).Table(table).Get(id).Delete().Exec(db.Session)
}

func (db *Database) GetElementById(table string, id string) (i interface{}, e error) {
	var cursor *r.Cursor

	if cursor, e = r.DB(db.Name).Table(table).Get(id).Run(db.Session); e != nil {
		return
	}

	if e = cursor.One(&i); e != nil {
		return
	}

	return
}

func (db *Database) GetElementsBy(table string, filter map[string]interface{}) (i interface{}, e error) {
	var cursor *r.Cursor

	if cursor, e = r.DB(db.Name).Table(table).Filter(filter).Run(db.Session); e != nil {
		return
	}

	if e = cursor.All(&i); e != nil {
		return
	}

	return
}
