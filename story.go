package cyoa

type Option struct {
	Chapter string `json:"arc"`
	Text    string `json:"text"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Story map[string]Chapter
