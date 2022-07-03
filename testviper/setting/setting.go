package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 对应配置文件编写结构体
type AppConfig struct {

	// mapstructure 将配置文件中的数据和结构体中的属性进行映射绑定
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int    `mapstructure:"machine_id"`

	// 通过指针引用获取其他结构体的结构并进行绑定
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// 创建结构体对象
var Conf AppConfig

// 通过 Init 函数加载配置文件并绑定到结构体对象中
func Init() error {
	// 按照指定路径读取配置文件
	viper.SetConfigFile("./config/config.yaml")

	// 实时监听和查看配置文件的变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改!!!")
		// 配置文件产生变化后重新解构赋值给结构体对象
		err := viper.Unmarshal(&Conf)
		if err != nil {
			return
		}
	})

	// 判断是否成功读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}

	// 将配置文件解构赋值给结构体对象
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return err
}
