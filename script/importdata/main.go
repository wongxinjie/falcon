package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rs/xid"

	"falcon/infra"
	"falcon/mapper/es/poetrymapper"
	"falcon/model/es/poetry"
)

func parseCSV(m *poetrymapper.Mapper, path string) ([]*poetry.Schema, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file error=%w", err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv file error=%w", err)
	}
	
	log.Printf("process data count: %d", len(rows))

	schemas := make([]*poetry.Schema, 0)
	for _, r := range rows {
		id := xid.New().String()
		s := &poetry.Schema{
			Id:        id,
			Title:     r[0],
			Dynasty:   r[1],
			Author:    r[2],
			Content:   r[3],
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		fmt.Println(*s)
		schemas = append(schemas, s)
	}
	return schemas, nil
}


func main() {
	client, err := infra.InitESClient()
	if err != nil {
		log.Printf("init esu client error=%v", err)
		return
	}

	m  := poetrymapper.New(client)
	rows, err := parseCSV(m, "/home/wongxinjie/Downloads/Poetry-master/唐.csv")
	if err != nil {
		log.Printf("parse csv file error=%v", err)
		return
	}

	rows = rows[:2000]
	err = m.BulkCreate(rows)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ok")

}
