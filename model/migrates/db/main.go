package main

import (
	"falcon/infra"
	"falcon/instance/loginst"
	"falcon/model/db/dynasty"
	"falcon/model/db/loginaccout"
)

var tables = []interface{}{
	dynasty.Schema{},
	loginaccout.Schema{},
}

func main() {
	db, err := infra.InitDB()
	if err != nil {
		loginst.Inst().Info("create mysql connection error")
		return
	}

	err = db.AutoMigrate(tables...)
	if err != nil {
		loginst.Inst().Info("auto migrate error")
	}
}
