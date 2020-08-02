package main

import "fmt"

func main() {
	// slice和基本操作和数组一样
	a := make([]string, 3)
	fmt.Println(a)

	a[1] = "hahahah"
	fmt.Println(a)
	fmt.Println(len(a))

	// slice的长度不定
	// append可以直接向slice中添加元素
	a = append(a, "gg")
	a = append(a, "ff")
	a[0] = "qqq"
	fmt.Println(a)

	// 复制操作
	b := make([]string, len(a))
	copy(b, a)
	fmt.Println(b)

	// 切片操作,类似substring
	a[2] = "2"
	fmt.Println(a)
	// 1-2
	c := a[1:3]
	fmt.Println(c)
	// 0-2
	c = a[:3]
	fmt.Println(c)
	// 3-last
	c = a[3:]
	fmt.Println(c)

	// 声明并初始化
	d := []string{"1", "2"}
	fmt.Println(d)

	e := make([][]string, 3)
	fmt.Println(e)
}
