package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	// 设置为生产环境
	gin.SetMode(gin.ReleaseMode)

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()
	// 强制日志颜色化
	//gin.ForceConsoleColor()

	// 记录到文件。
	f, _ := os.Create("./gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 用默认中间件创建一个 gin 路由:
	// 日志和恢复（无崩溃）中间件
	r := gin.Default()

	// 禁用可信代理功能
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
