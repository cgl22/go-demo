package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "str1"
	}()
	select {
	case msg := <-c1:
		fmt.Println(msg)
	case a := <-time.After(time.Second):
		fmt.Println(a)
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		c2 <- "str2"
	}()
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout2")
	}
}
