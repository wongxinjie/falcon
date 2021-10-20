package poetry

const (
	Index    = "poetry"
	Mappging = `{
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword"
				},
				"title": {
					"type": "keyword"
				},
				"dynastymapper": {
					"type": "keyword"
				},
				"content": {
					"type": "text",
					"analyzer": "ik_smart",
					"search_analyzer": "ik_smart"
				},
				"created_at": {
					"type": "long"
				},
				"updated_at": {
					"type": "long"
				}
			}
		}
	}`
)

type Schema struct {
	indexName struct{} `json:"poetry"`
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Dynasty   string   `json:"dynastymapper"`
	Author    string   `json:"author"`
	Content   string   `json:"content"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}
