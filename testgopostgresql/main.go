package main

import (
	"fmt"
	"testgopostgresql/setting"
)

func main() {
	fmt.Println("hello world!!!")
	setting.Init()
	fmt.Println(setting.SiteInfo)
}
