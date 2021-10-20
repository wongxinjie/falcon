package poetrymapper

import (
	"context"
	"fmt"
	"testing"

	"falcon/infra"
)

func TestMapper_SearchByWord(t *testing.T) {
	client, err := infra.InitESClient()
	if err != nil {
		t.Errorf("error=%v", err)
		return
	}

	m := New(client)

	fmt.Println(m.Tag(&m.Dynasty).V(), m.IndexName())

	count, rows, err := m.SearchByWord(context.Background(), "秋", 0, 10)
	if err != nil {
		t.Errorf("error=%v", err)
		return
	}
	t.Log(count)

	for _, r := range rows {
		t.Logf("row=%v", r)
	}
}

func TestMapper_DropIndex(t *testing.T) {
	client, err := infra.InitESClient()
	if err != nil {
		t.Errorf("error=%v", err)
		return
	}

	m := New(client)
	err = m.DropIndex(context.Background())
	if err != nil {
		t.Errorf("error=%v", err)
		return
	}
}

func TestMapper_Detail(t *testing.T) {
	client, err := infra.InitESClient()
	if err != nil {
		t.Errorf("error=%v", err)
		return
	}

	m := New(client)

	p, err := m.Detail(context.Background(), "13066109802680197")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", p)
	}

}
