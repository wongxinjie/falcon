package config

import "fmt"

type ESConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (es *ESConfig) Url() string {
	return fmt.Sprintf("http://%s:%d", es.Host, es.Port)
}

func (es *ESConfig) IsValid() bool {
	return true
}
