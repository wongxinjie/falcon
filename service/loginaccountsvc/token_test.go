package loginaccountsvc

import (
	"context"
	"testing"
	"time"

	"falcon/enum"
	"falcon/infra"
)

func TestCacheLoginUser(t *testing.T) {
	ifr, err := infra.GetTestInfra()
	if err != nil {
		t.Fatal(err)
	}

	userData := &LoginUserCacheData{
		UserID:  10,
		LoginAt: time.Now().UTC().Unix(),
		Via:     enum.LoginViaWeb,
		Status:  0,
	}

	err = CacheLoginUser(context.TODO(), ifr, userData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestFetchCachedLoginUser(t *testing.T) {
	ifr, err := infra.GetTestInfra()

	if err != nil {
		t.Fatal(err)
	}

	user, err := FetchCachedLoginUser(context.TODO(), ifr, 10, enum.LoginViaWeb)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", user)
}
