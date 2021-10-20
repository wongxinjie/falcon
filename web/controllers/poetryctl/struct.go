package poetryctl

type PoetryData struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SearchResponse struct {
	Total  int64         `json:"total"`
	Poetry []*PoetryData `json:"poetry"`
}

type DetailResponse struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Dynasty string `json:"dynasty"`
	Content string `json:"content"`
}
