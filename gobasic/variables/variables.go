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
	var number int64 = 18
	fmt.Println(testFun(number))
	fmt.Println(Mytype)
	// 简短变量声明，类型自动推导
	n := 55
	fmt.Println(testFun(int64(n)))
}

/*
	func 函数名 (参数列表) 返回值{
		 函数体
		 ...
	}

*/
func testFun(n int64) int64 {
	return n
}
