package loginaccountsvc

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"

	"falcon/enum/redisenum"
	"falcon/infra"
)

type LoginUserCacheData struct {
	UserID  int64  `json:"user_id"`
	LoginAt int64  `json:"login_at"`
	Via     string `json:"via"`
	Status  int    `json:"status"`
}

func CacheLoginUser(ctx context.Context, ifr *infra.Infra, user *LoginUserCacheData) error {
	rdb := ifr.RedisDB
	key := redisenum.GetLoginUserCacheKey(user.UserID)

	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	_, err = rdb.HSetNX(ctx, key, user.Via, string(payload)).Result()
	if err != nil {
		return err
	}

	return nil
}

func FetchCachedLoginUser(ctx context.Context, ifr *infra.Infra, userID int64, via string) (*LoginUserCacheData, error) {
	rdb := ifr.RedisDB
	key := redisenum.GetLoginUserCacheKey(userID)

	payload, err := rdb.HGet(ctx, key, via).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(payload) == 0 {
		return new(LoginUserCacheData), nil
	}

	var user LoginUserCacheData
	err = json.Unmarshal([]byte(payload), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
