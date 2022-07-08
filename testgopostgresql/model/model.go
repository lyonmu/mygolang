package model

import "fmt"

type Site struct {
	Id         int    `gorm:"primary_key" json:"id"`
	Logo       string `json:"username"`
	Icon       string `json:"nickname"`
	Title      string `json:"email"`
	Slogan     string `json:"avatar"`
	CreateTime string `json:"description"`
	BackGround string `json:"title"`
	Ministry   string `json:"miit"`
	Police     string `json:"author"`
}

func Init() {
	fmt.Println("model")
}
