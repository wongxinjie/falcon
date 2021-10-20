package token

import (
	"testing"

	"falcon/config"
	"falcon/enum"
)

func TestGenerateToken1(t *testing.T) {
	token, err := GenerateToken(10, enum.LoginViaWeb, []byte(config.UserJwtSecret))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	token, err := GenerateToken(10, enum.LoginViaWeb, []byte(config.UserJwtSecret))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	ok, u := ParseToken(token, []byte(config.UserJwtSecret))
	if !ok {
		t.Log(ok)
	} else {
		t.Logf("%+v", u)
	}

}
