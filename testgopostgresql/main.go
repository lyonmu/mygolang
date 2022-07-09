package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testgopostgresql/setting"
)

func main() {
	fmt.Println("hello world!!!")
	setting.Init()
	// 普通输出结构体对象信息
	// fmt.Println(setting.SiteInfo)

	// JSON 格式输出结构体对象信息
	data, err := json.MarshalIndent(setting.SiteInfo, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
