package database

const (
	COMMUNITY = "community" // Created by the community, but verified by the Aridia team.
	OFFICIAL  = "official"  // Created by the Aridia team.
	NONE      = "none"      // Created by the community, but not verified by the Aridia team.
)

type Step struct {
	Name string `json:"name"`
	Desc string `json:"desc"`

	ImageURL string `json:"image_url"`
	Number   string `json:"number"`
}

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`

	Answer string `json:"answer"`
	Check  string `json:"check"`
	Steps  []Step `json:"steps"`

	CreatedAt string `json:"created_at"`
	AuthorID  string `json:"author_id"`
}
