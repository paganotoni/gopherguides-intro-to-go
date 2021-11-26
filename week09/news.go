package week09

// News represents a news item, the core concept of our system.
type News struct {
	ID int `json:"id"`

	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Source  string `json:"source"`

	Categories []string `json:"category"`
}
