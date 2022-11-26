/*
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
*/

const resp = await fetch('http://localhost:8000/api/problems', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    },
    credentials: 'same-origin',
    body: JSON.stringify({
        title: 'test',
		description: 'test',
		explanation: [],
		author_id: '347787075308748801',
		answer: 'test',
    }),
});

await resp.json();
