package esu

import "github.com/olivere/elastic/v7"

type Mapper struct {
	*elastic.Client
	*Meta
}

func New(c *elastic.Client, meta *Meta) *Mapper {
	m := &Mapper{
		Client: c,
		Meta:   meta,
	}
	return m
}
