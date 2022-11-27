package database

import "golang.org/x/oauth2"

type User struct {
	Id    string `json:"id"`
	Admin bool   `json:"admin"`

	Auth oauth2.Token `json:"auth"`

	Badges []string `json:"badges"`
	Posts  []string `json:"posts"`
}
