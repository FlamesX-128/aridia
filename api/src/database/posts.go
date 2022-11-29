package database

import (
	"github.com/FlamesX-128/aridia/api/src/models/database"
)

func (db *Database) UpdatePost(post *database.Post) error {
	return db.UpdateElement("posts", post.Id, post)
}

func (db *Database) CreatePost(post *database.Post) error {
	return db.CreateElement("posts", post)
}

func (db *Database) DeletePost(id string) error {
	return db.DeleteElement("posts", id)
}

func (db *Database) GetPostsBy(filter map[string]interface{}) ([]database.Post, error) {
	i, e := db.GetElementsBy("posts", filter)

	return i.([]database.Post), e
}

func (db *Database) GetPostById(id string) (database.Post, error) {
	i, e := db.GetElementById("posts", id)

	return i.(database.Post), e
}
