package main

import "fmt"

func main() {
	f("单线程")

	// 使用go启动协程
	go f("线程1")
	go f("线程2")

	// 为匿名函数使用协程
	go func(str string) {
		fmt.Println(str)
	}("线程-匿名函数")

	// 监听输入，阻塞主线程
	var input string
	fmt.Scanln(&input)
	fmt.Println("主线程结束")
}

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, "：", i)
	}
}
