package loginaccountmapper

import (
	"context"
	"testing"

	"falcon/infra"
)

func TestMapper_CreateAccount(t *testing.T) {
	ifr, err := infra.GetTestInfra()
	if err != nil {
		t.Fatal(err)
	}

	m := New(ifr.DB)
	ok, err := m.CreateAccount(context.Background(), "18825111143", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ok)

	account, err := m.PasswordLogin(context.Background(), "18825111143", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", account)

}
