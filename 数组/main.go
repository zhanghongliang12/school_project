package main

import "fmt"

func main() {
	var s [3][5]int

	fmt.Println(s)
	array := [...]int{1, 3, 4, 5, 6, 7}
	fmt.Println(array)
	fmt.Println(array[4])
	array[1] = 9999 // 替换元素
	fmt.Println(array)
	// 循环打印
	// 打印
	for i, v := range array {
		fmt.Println(i, v)
	}
	// 循环
	for i := 0; i < len(array); i++ {
		fmt.Println(i, array[i])
	}

}
