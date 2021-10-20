package config

import "fmt"

const (
	charset   = "utf8mb4"
	parseTime = true
	location  = ""
)

type DBConfig struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	User               string `json:"user"`
	Password           string `json:"password"`
	DB                 string `json:"db"`
	MaxIdleConn        int    `json:"max_idle_conn"`
	MaxOpenConn        int    `json:"max_open_conn"`
	MaxIdleTimeSeconds int    `json:"max_idle_time_seconds"`
	MaxLifeTimeSeconds int    `json:"max_life_time_seconds"`
}

func (db *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		db.User, db.Password, db.Host, db.Port, db.DB, charset, parseTime, location)
}

func (db *DBConfig) IsValid() bool {
	return len(db.Host) > 0 && db.Port > 0
}
