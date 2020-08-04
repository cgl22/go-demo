package main

import "fmt"

func main() {
	c1 := make(chan string, 3)
	c1 <- "s1"
	c1 <- "s2"
	c1 <- "s3"
	<-c1
	c1 <- "s4"
	close(c1)

	// 遍历channel中的值
	for str := range c1 {
		fmt.Println(str)
	}
}
