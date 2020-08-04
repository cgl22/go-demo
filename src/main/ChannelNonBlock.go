package main

import "fmt"

func main() {
	c1 := make(chan string, 1)
	// 使用default实现非阻塞通道通信
	select {
	case msg := <-c1:
		fmt.Println(msg)
	default:
		fmt.Println("no msg")
	}
}
