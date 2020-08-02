package main

import "fmt"

func main() {
	var a string = "ss"
	fmt.Println(a)

	// 可以声明多个
	var b, c int = 3, 4
	fmt.Println(b, c)

	// 自动类型推断
	var d, e = 5.0, "hahaha"
	fmt.Println(d, e)

	// 简写
	f := true
	fmt.Println(f)

	// 默认赋0值
	var g float32
	fmt.Println(g)
	var h bool
	fmt.Println(h)
}
