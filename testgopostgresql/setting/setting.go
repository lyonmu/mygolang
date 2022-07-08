package setting

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"testgopostgresql/model"
)

func Init() {

	fmt.Println("setting package")
	dsn := "host=localhost user=muqing password= database=muqing port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info), // 打印执行的 SQL 语句
	})
	if err != nil {
		fmt.Println("数据库连接出现错误")
		fmt.Println(err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库连接池出现错误")
		fmt.Println(err)
		return
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	model.Init()

	var site model.Site

	db.First(&site)

	fmt.Println(site)

}
