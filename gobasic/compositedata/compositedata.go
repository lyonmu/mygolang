package compositedata

import (
	"fmt"
)

func Init() {
	fmt.Println(Description)
	testArr()
	testSlice()
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

func testSlice() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	arr := [...]int{1, 2, 3, 4, 5}
	reverse(arr[:])
	fmt.Println(arr)

	fmt.Println(len(arr))
	fmt.Println(len(summer))
	fmt.Println(cap(arr))
	fmt.Println(cap(Q2))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
