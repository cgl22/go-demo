package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		c1 <- "str1"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "str2"
	}()

	// select同时等待c1和c2，若满足case则运行完成
	// 循环两次，获取两次
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
