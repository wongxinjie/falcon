package dbu

import (
	"gorm.io/gorm"
)

type Mapper struct {
	*gorm.DB
	*Meta
}

func New(c *gorm.DB, meta *Meta) *Mapper {
	m := &Mapper{
		DB:   c,
		Meta: meta,
	}
	return m
}
