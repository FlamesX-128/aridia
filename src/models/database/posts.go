package database

const (
	COMMUNITY = "community" // Created by the community, but verified by the Aridia team.
	OFFICIAL  = "official"  // Created by the Aridia team.
	NONE      = "none"      // Created by the community, but not verified by the Aridia team.
)

type Step struct {
	Name string `json:"name" rethinkdb:"name,omitempty"`
	Desc string `json:"desc" rethinkdb:"desc"`

	ImageURL string `json:"image_url" rethinkdb:"image_url"`
	Number   string `json:"number" rethinkdb:"number,omitempty"`
}

type Post struct {
	Id    string `json:"id" rethinkdb:"id,omitempty"`
	Title string `json:"title" rethinkdb:"title,omitempty"`
	Desc  string `json:"desc" rethinkdb:"desc,omitempty"`

	Answer string `json:"answer" rethinkdb:"answer,omitempty"`
	Check  string `json:"check" rethinkdb:"check,omitempty"`
	Steps  []Step `json:"steps" rethinkdb:"steps"`

	CreatedAt string `json:"created_at" rethinkdb:"created_at,omitempty"`
	AuthorID  string `json:"author_id" rethinkdb:"author_id,omitempty"`
}
