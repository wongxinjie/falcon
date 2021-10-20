package config

type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func (rdb *RedisConf) IsValid() bool {
	return len(rdb.Addr) > 0
}
