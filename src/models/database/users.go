package database

import "golang.org/x/oauth2"

type SimplifiedUser struct {
	Id string `json:"id"`

	Badges []string `json:"badges"`
	Posts  []string `json:"posts"`
}

type User struct {
	Id    string `json:"id" rethinkdb:"id,omitempty"`
	Admin bool   `json:"admin" rethinkdb:"admin"`

	Auth oauth2.Token `json:"auth" rethinkdb:"auth,omitempty"`

	Badges []string `json:"badges" rethinkdb:"badges"`

	LastPost string `json:"last_post" rethinkdb:"last_post"`
}
