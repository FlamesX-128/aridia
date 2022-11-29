package database

type Content struct {
	ImageUrl string `json:"image_url" rethinkdb:"image_url,omitempty"`
	Desc     string `json:"desc" rethinkdb:"desc,omitempty"`
	Name     string `json:"name" rethinkdb:"name,omitempty"`
}

type Solution struct {
	Name    string    `json:"name" rethinkdb:"name,omitempty"`
	Content []Content `json:"content" rethinkdb:"content"`
}

type Post struct {
	AuthorId string `json:"author_id" rethinkdb:"author_id,omitempty"`
	Id       string `json:"id" rethinkdb:"id,omitempty"`

	Title    string `json:"title" rethinkdb:"title,omitempty"`
	Desc     string `json:"desc" rethinkdb:"desc"`
	ImageUrl string `json:"image_url" rethinkdb:"image_url"`

	Solutions []Solution `json:"solution" rethinkdb:"solution"`
	Answers   []string   `json:"answers" rethinkdb:"answers"`

	CreatedAt string `json:"created_at" rethinkdb:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at" rethinkdb:"updated_at,omitempty"`
}
