package main

import "fmt"

func main() {
	// 数组默认是零值
	var a [5]int
	fmt.Println(a)

	a[2] = 99
	fmt.Println(a[0])
	fmt.Println(a[2])
	fmt.Println(len(a))

	// 声明并初始化,b[4]为零值
	b := [5]int{1, 2, 3, 4}
	fmt.Println(b)

	// 二维数组
	var c [2][3]int
	fmt.Println(c)
}
