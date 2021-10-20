package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("/tmp/falcon.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(viper.Get("db.host"))

	return

}
