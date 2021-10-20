package poetrymapper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"

	"falcon/model/es/poetry"
	"falcon/pkg/meta/esu"
)

var (
	S poetry.Schema
	m esu.Meta
)

func init() {
	m.Init(&S)
}

type Mapper struct {
	*esu.Mapper
	*poetry.Schema
}

func New(c *elastic.Client) *Mapper {
	return &Mapper{
		Mapper: esu.New(c, &m),
		Schema: &S,
	}
}

func (m *Mapper) BulkCreate(rows []*poetry.Schema) error {
	req := m.Client.Bulk().Index(m.IndexName())
	for _, r := range rows {
		doc := elastic.NewBulkIndexRequest().
			Id(r.Id).
			Doc(r)
		req.Add(doc)
	}

	if req.NumberOfActions() < 0 {
		return nil
	}

	_, err := req.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (m *Mapper) SearchByWord(ctx context.Context, word string, offset, limit int) (int64, []*poetry.Schema, error) {
	match := elastic.NewMatchQuery(
		m.Tag(&m.Content).V(),
		word)

	searchResult, err := m.Client.Search().
		Index(m.IndexName()).
		Query(match).
		From(offset).
		Sort(m.Tag(&m.Id).V(), true).
		Size(limit).
		Do(ctx)
	if err != nil {
		return 0, nil, fmt.Errorf("search error=%w", err)
	}

	if searchResult.TotalHits() == 0 {
		return 0, make([]*poetry.Schema, 0), nil
	}

	total := searchResult.TotalHits()
	results := make([]*poetry.Schema, 0)
	for _, hit := range searchResult.Hits.Hits {
		var p poetry.Schema
		if err := json.Unmarshal(hit.Source, &p); err != nil {
			return 0, nil, fmt.Errorf("parse result error=%w, source=%+v", err, hit.Source)
		}
		results = append(results, &p)
	}

	return total, results, nil
}

type MultiQuery struct {
	Title   string
	Author  string
	Dynasty string
	Content string
}

func (m *Mapper) Search(ctx context.Context, q *MultiQuery, limit, offset int) ([]*poetry.Schema, error) {
	return nil, nil
}

func (m *Mapper) Detail(ctx context.Context, id string) (*poetry.Schema, error) {
	result, err := m.Client.Get().
		Index(m.IndexName()).
		Id(id).
		Do(ctx)
	if err != nil && !elastic.IsNotFound(err) {
		return nil, err
	}
	if elastic.IsNotFound(err) {
		return &poetry.Schema{}, nil
	}

	var p poetry.Schema
	if err := json.Unmarshal(result.Source, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (m *Mapper) DropIndex(ctx context.Context) error {
	_, err := m.Client.DeleteIndex(m.IndexName()).Do(ctx)
	if err != nil {
		return fmt.Errorf("delete index error=%w, index=%s", err, m.IndexName())
	}

	return nil
}
