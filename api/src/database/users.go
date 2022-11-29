package database

import (
	"github.com/FlamesX-128/aridia/api/src/models/database"
)

func (db *Database) UpdateUser(user *database.User) error {
	return db.UpdateElement("users", user.Id, user)
}

func (db *Database) CreateUser(user *database.User) error {
	return db.CreateElement("users", user)
}

func (db *Database) DeleteUser(id string) error {
	return db.DeleteElement("users", id)
}

func (db *Database) GetUsersBy(filter map[string]interface{}) ([]database.User, error) {
	i, e := db.GetElementsBy("users", filter)

	return i.([]database.User), e
}

func (db *Database) GetUserById(id string) (database.User, error) {
	i, e := db.GetElementById("users", id)

	return i.(database.User), e
}
