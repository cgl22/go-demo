package main

import "fmt"

const ss string = "constA"

func main() {
	fmt.Println(ss)

	// 常数表达式可以执行任意精度的运算
	const b = 400
	const c = 3e20 / b
	fmt.Println(c)

	// 数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化
	fmt.Println(int64(c))
}
