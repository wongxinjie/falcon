package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"

	"falcon/config"
)

func NewESClient(c *config.ESConfig) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(c.Url()),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		return nil, fmt.Errorf("new elasticsearch client error=%w", err)
	}
	return client, nil
}

func InitESClient() (*elastic.Client, error) {
	conf := &config.ESConfig{
		Host: "127.0.0.1",
		Port: 9200,
	}

	return NewESClient(conf)
}
