package poetrysrv

import (
	"context"

	"falcon/infra"
	"falcon/mapper/es/poetrymapper"
	"falcon/model/es/poetry"
)

type PoetryService struct {
	ifr *infra.Infra
	poetryMapper *poetrymapper.Mapper
}

const (
	TypeTitle   = 1
	TypeAuthor  = 2
	TypeDynasty = 3
	TypeContent = 4
)

type SearchArgs struct {
	Type  int
	Query string
}

func New(ifr *infra.Infra) *PoetryService {
	return &PoetryService{
		ifr: ifr,
		poetryMapper: poetrymapper.New(ifr.ESClient),
	}
}

func (s *PoetryService) Search(ctx context.Context, q *SearchArgs, offset, limit int) (int64, []*poetry.Schema, error) {

	switch q.Type {
	case TypeContent:
		return s.poetryMapper.SearchByWord(ctx, q.Query, offset, limit)
	}

	return 0, make([]*poetry.Schema, 0), nil
}

func (s *PoetryService) Detail(ctx context.Context, id string) (*poetry.Schema, error) {

	return s.poetryMapper.Detail(ctx, id)
}
