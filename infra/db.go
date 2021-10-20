package infra

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"falcon/config"
	"falcon/instance/loginst"
	"falcon/pkg/logging"
)

func NewDB(c *config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.Open(c.DSN()), &gorm.Config{
			Logger: logging.NewGormLogger(*loginst.Inst(), time.Duration(1)*time.Second),
		})
	if err != nil {
		return nil, fmt.Errorf("NewDB error=%w, dsn=%s", err, c.DSN())
	}

	innerDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error=%w", err)
	}
	if c.MaxIdleConn > 0 {
		innerDB.SetMaxIdleConns(c.MaxIdleConn)
	}
	if c.MaxOpenConn > 0 {
		innerDB.SetMaxOpenConns(c.MaxOpenConn)
	}

	if c.MaxIdleTimeSeconds > 0 {
		innerDB.SetConnMaxIdleTime(time.Duration(c.MaxIdleTimeSeconds) * time.Second)
	}
	if c.MaxLifeTimeSeconds > 0 {
		innerDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTimeSeconds) * time.Second)
	}
	return db, err
}
