/*


type Step struct {
	Name string `json:"name"`
	Desc string `json:"desc"`

	ImageURL string `json:"image_url"`
	Number   string `json:"number"`
}

Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`

	Answer string `json:"answer"`
	Check  string `json:"check"`
	Steps  []Step `json:"steps"`

	CreatedAt string `json:"created_at"`
	AuthorID  string `json:"author_id"`

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
		desc: 'test',
		answer: 'test',
		stepts: [
			{ name: 'test', desc: 'test', number: "1" },
		]
	}),
});

await resp.json();
