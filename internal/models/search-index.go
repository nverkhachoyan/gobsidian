package models

type SearchIndex struct {
	Notes []*SearchIndexNote `json:"notes"`
}

type SearchIndexNote struct {
	ID    int64    `json:"id"`
	Title string   `json:"title"`
	URL   string   `json:"url"`
	IMG   string   `json:"img"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}
