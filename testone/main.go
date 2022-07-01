package main

import (
	"fmt"
	"github.com/6tail/lunar-go/calendar"
)

func main() {
	lunar := calendar.NewLunarFromYmd(1998, 8, 11)
	fmt.Println(lunar.ToFullString())
	fmt.Println(lunar.GetSolar().ToFullString())
}
