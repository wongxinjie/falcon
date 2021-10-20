package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	UserJwtSecret = "123456789"
)

var (
	dbConf    *DBConfig
	redisConf *RedisConf
	esConf    *ESConfig
)

type loadConfigFunc func() error

func LoadConfig() error {
	confPath := os.Getenv("falcon_path")
	if len(confPath) == 0 {
		confPath = "/tmp/falcon.yaml"
	}

	viper.SetConfigFile(confPath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	applyFns := []loadConfigFunc{
		initDBConf,
		initESConf,
		initRedisConf,
	}

	for _, fn := range applyFns {
		if err := fn(); err != nil {
			panic(fmt.Errorf("LoadConfig error=%+v", err))
		}
	}
	return nil
}

func AppDBConfig() *DBConfig {
	return dbConf
}

func AppRedisConfig() *RedisConf {
	return redisConf
}

func AppESConfig() *ESConfig {
	return esConf
}

func initDBConf() error {
	dbConf = &DBConfig{
		Host:               viper.GetString("db.host"),
		Port:               viper.GetInt("db.port"),
		User:               viper.GetString("db.user"),
		Password:           viper.GetString("db.password"),
		DB:                 viper.GetString("db.db"),
		MaxIdleConn:        viper.GetInt("db.max_idle_conn"),
		MaxOpenConn:        viper.GetInt("db.max_open_conn"),
		MaxIdleTimeSeconds: viper.GetInt("db.max_idle_time_seconds"),
		MaxLifeTimeSeconds: viper.GetInt("db.max_lift_time_seconds"),
	}

	if !dbConf.IsValid() {
		return fmt.Errorf("db config is invalid, conf=%+v", dbConf)
	}
	return nil
}

func initRedisConf() error {
	redisConf = &RedisConf{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
	if !redisConf.IsValid() {
		return fmt.Errorf("redis config is invalid, conf=%+v", redisConf)
	}
	return nil
}

func initESConf() error {
	esConf = &ESConfig{
		Host: viper.GetString("es.host"),
		Port: viper.GetInt("es.port"),
	}
	if !esConf.IsValid() {
		return fmt.Errorf("es config is invalid, conf=%+v", esConf)
	}
	return nil
}
