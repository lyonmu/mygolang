package variables

import "fmt"

/*
	type 类型名 struct{
		 属性名 属性类型
		 ...
	}
*/

type TestType struct {
	Id        int64
	Name      string
	IsStudent bool
}

var Mytype = TestType{
	1921681510, "lyonmu", true,
}

func Init() {
	fmt.Println(Description)
	fmt.Println("=========================")
	var number int = 18
	// 简短变量声明，类型自动推导
	n := 55
	fmt.Println(testFun(number))
	fmt.Println(Mytype)
	fmt.Println("=========================")
	fmt.Println(testFun(n))
	fmt.Println(n)
	fmt.Println("=========================")
	testPointer(&n)
	fmt.Println(n)

	testvar := new(int)
	fmt.Println(testvar) // 打印变量的内存地址
	fmt.Println(*testvar)
}

/*
	func 函数名 (参数列表) 返回值{
		 函数体
		 ...
	}

*/
func testFun(n int) int {
	n = 72
	return n
}

func testPointer(p *int) {
	*p = 89
}
