package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := make(chan int, 1)
	done := make(chan bool, 1)

	go func() {
		for {
			// tasks关闭后more才为false
			j, more := <-tasks
			if more {
				fmt.Println("receive:", j)
			} else {
				fmt.Println("receive all")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("send:", i)
		tasks <- i
		time.Sleep(time.Second)
	}
	fmt.Println("send all")
	close(tasks)

	<-done
}
