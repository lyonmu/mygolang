package flowcontrol

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

func Init() {
	testIf(-20)
	testFor()
	nine()
	testFR()
	testSwitch()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	logger.Info("hello!", zap.String("name", "xiaomin"), zap.Int("age", 20))
}

func testIf(n int) {
	if n == 0 {
		fmt.Println("n = 0")
	} else if n > 0 {
		fmt.Println("n > 0")
	} else {
		fmt.Println("n < 0")
	}
}

func testFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\t", i)
	}
	fmt.Println()
}

func nine() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func testFR() {
	arr := [...]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

func testSwitch() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
