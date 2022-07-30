package compositedata

import "fmt"

func Init() {
	fmt.Println(Description)
	testArr()
}

func testArr() {
	// 长度为 3 且数组里面值为 1 2 3 4 5
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("%d %d\n", i, v)
	}

	// 长度为 3 且数组里面值为 1 2 3
	a := [...]int{1, 2, 3}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 长度为 3 且最后一位为 -1 的数组
	r := [...]int{2: -1}
	for _, v := range r {
		fmt.Printf("%d\n", v)
	}

	testUpdate(&a)
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

}

// 通过指针修改数组值,函数的形参必须指定数组的长度
func testUpdate(arr *[3]int) {
	for i := range arr {
		arr[i] = 0
	}
}
