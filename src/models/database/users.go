package database

type User struct {
	Permisson uint8  `json:"permission"`
	ID        string `json:"id"`
}
