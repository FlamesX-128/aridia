package tools

import (
	mdb "github.com/FlamesX-128/aridia/src/models/database"
	"golang.org/x/oauth2"
)

func CreateUser(id string, auth oauth2.Token) mdb.User {
	return mdb.User{
		Id:       id,
		Admin:    false,
		Auth:     auth,
		Badges:   []string{},
		Solved:   []string{},
		LastPost: "0",
	}
}
