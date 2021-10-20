package infra

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"

	"falcon/config"
)

var (
	realInfra *Infra
)

type Infra struct {
	ESClient *elastic.Client
	RedisDB  *redis.Client
	DB       *gorm.DB
}

type wrapperInfra struct {
	ifr *Infra
	err error
}

func SetUp() *wrapperInfra {
	return &wrapperInfra{
		ifr: new(Infra),
		err: nil,
	}
}

func (c *wrapperInfra) WithDB() *wrapperInfra {
	if c.err != nil {
		return c
	}

	db, err := NewDB(config.AppDBConfig())
	if err != nil {
		c.err = err
		return c
	}

	c.ifr.DB = db
	return c
}

func (c *wrapperInfra) WithRedis() *wrapperInfra {
	if c.err != nil {
		return c
	}

	rdb, err := NewRedisClient(config.AppRedisConfig())
	if err != nil {
		c.err = err
		return c
	}

	c.ifr.RedisDB = rdb
	return c
}

func (c *wrapperInfra) WithES() *wrapperInfra {
	if c.err != nil {
		return c
	}

	es, err := NewESClient(config.AppESConfig())
	if err != nil {
		c.err = err
		return c
	}

	c.ifr.ESClient = es
	return c
}

func (c *wrapperInfra) Build() error {
	if c.err != nil {
		return c.err
	}

	realInfra = c.ifr
	return nil
}

func Inst() *Infra {
	return realInfra
}

func GetTestInfra() (*Infra, error) {
	err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	err = SetUp().
		WithDB().
		WithRedis().
		Build()
	if err != nil {
		return nil, err
	}

	return realInfra, nil
}
