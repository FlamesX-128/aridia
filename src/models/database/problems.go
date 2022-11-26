package database

type Explanation struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Number  uint8  `json:"number"`
}

type PostProblem struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Explanation []Explanation `json:"explanation"`
	AuthorID    string        `json:"author_id"`
	Answer      string        `json:"answer"`
}

type PutProblem struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Explanation []Explanation `json:"explanation"`
	AuthorID    string        `json:"author_id"`
	Answer      string        `json:"answer"`
}

type GetProblem struct {
	ID          string        `json:"id"`
	Description string        `json:"description"`
	Explanation []Explanation `json:"explanation"`
	AuthorID    string        `json:"author_id"`
	Answer      string        `json:"answer"`
	Community   bool          `json:"community"`
	Verified    bool          `json:"verified"`
	CreatedAt   string        `json:"created_at"`
}

type Problem struct {
	Description string        `json:"description"`
	Explanation []Explanation `json:"explanation"`
	AuthorID    string        `json:"author_id"`
	Answer      string        `json:"answer"`
	Community   bool          `json:"community"`
	Verified    bool          `json:"verified"`
	CreatedAt   string        `json:"created_at"`
}
