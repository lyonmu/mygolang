package main

import (
	"fmt"
	settings "testviper/setting"
	jsonconfig "testviper/testjsonconfig"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	fmt.Println(settings.Conf.RedisConfig)

	if err := jsonconfig.Init(); err != nil {
		fmt.Printf("load json config failed, err:%v\n", err)
		return
	}
	fmt.Println(jsonconfig.InfoConf)
}
