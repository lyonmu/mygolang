package main

import (
	"fmt"
	settings "testviper/setting"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	fmt.Println(settings.Conf.MySQLConfig)
}
