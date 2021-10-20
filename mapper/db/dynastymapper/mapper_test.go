package dynastymapper

import (
	"context"
	"testing"

	"falcon/infra"
	"falcon/instance/loginst"
	"falcon/model/db/dynasty"
)

func TestMapper_Insert(t *testing.T) {
	db, err := infra.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	m := New(db)
	row := &dynasty.Schema{
		Name:  "宋",
		Count: 2000,
	}

	err = m.Insert(context.Background(), row)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("row=%+v", row)
}

func TestMapper_OneByName(t *testing.T) {
	db, err := infra.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	var name = "宋"
	m := New(db)
	row, err := m.OneByName(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("row=%+v", row)
}

func TestMapper_Update(t *testing.T) {
	db, err := infra.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	var name = "宋"
	m := New(db)
	row, err := m.OneByName(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}
	row.Count = 2000
	err = m.Update(context.Background(), row)
	if err != nil {
		t.Fatal(err)
	}

	loginst.Inst().Info("ok ")

}
