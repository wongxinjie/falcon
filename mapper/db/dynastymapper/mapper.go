package dynastymapper

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"falcon/model/db/dynasty"
	"falcon/pkg/meta/dbu"
)

var (
	S dynasty.Schema
	m dbu.Meta
)

func init() {
	m.Init(&S)
}

type Mapper struct {
	*dbu.Mapper
	*dynasty.Schema
}

func New(db *gorm.DB) *Mapper {
	return &Mapper{
		Mapper: dbu.New(db, &m),
		Schema: &S,
	}
}

func (m *Mapper) Insert(ctx context.Context, rows ...*dynasty.Schema) error {
	result := m.DB.Create(&rows)
	if result.Error != nil {
		return fmt.Errorf("insert error=%w", result.Error)
	}
	return nil
}

func (m *Mapper) OneByName(ctx context.Context, name string) (*dynasty.Schema, error) {
	var row dynasty.Schema
	result := m.DB.
		Where(m.Tag(&m.Schema.Name).Eq(), name).
		Take(&row)
	if result.Error != nil {
		return nil, fmt.Errorf("OneByName error=%w", result.Error)
	}

	return &row, nil
}

func (m *Mapper) Update(ctx context.Context, row *dynasty.Schema) error {
	result := m.DB.Model(row).Update(m.Tag(&m.Schema.Count).V(), row.Count)
	if result.Error != nil {
		return fmt.Errorf("update error=%w", result.Error)
	}
	return nil
}
