package main

import (
	"fmt"
	"time"
)

// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
// 这个特性允许我们，不使用任何其它的同步操作，来在程序结尾等待消息。
func main() {
	message := make(chan string)

	go func() { message <- "str1" }()
	str1 := <-message
	fmt.Println(str1)

	go func() { message <- "str2" }()
	str2 := <-message
	fmt.Println(str2)

	message2 := make(chan string, 2)
	message2 <- "str3"
	message2 <- "str4"
	str3 := <-message2
	str4 := <-message2
	fmt.Println(str3)
	fmt.Println(str4)

	done := make(chan bool)
	// 开启协程运行work方法
	go work(done)
	// 直到done发出通知前将一直阻塞
	<-done
}

func work(done chan bool) {
	fmt.Println("start")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
