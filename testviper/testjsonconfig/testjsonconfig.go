package testjsonconfig

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Info struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Miit        string `json:"miit"`
	Author      string `json:"author"`
	CreateTime  int64  `json:"createTime"`
}

var InfoConf Info

func Init() error {
	viper.SetConfigFile("./config/info.json")

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改!!!")
		err := viper.Unmarshal(&InfoConf)
		if err != nil {
			return
		}
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}
	if err := viper.Unmarshal(&InfoConf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
