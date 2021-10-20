package main

import (
	"context"
	"log"

	"falcon/infra"
	"falcon/model/es/poetry"
)

var mappings = []string{
	poetry.Mappging,
}

var indexes = []string{
	poetry.Index,
}


func main() {
	client, err := infra.InitESClient()
	if err != nil {
		log.Printf("create esu client error=%v", err)
		return
	}

	ctx := context.Background()
	for i, index := range indexes {
		ok, err := client.IndexExists(index).Do(ctx)
		if err != nil {
			log.Printf("query index error=%v", err)
			return
		}
		if !ok {
			_, err :=  client.CreateIndex(index).Body(mappings[i]).Do(ctx)
			if err != nil {
				log.Printf("create index error=%v", err)
				return
			}
		}
	}
}
