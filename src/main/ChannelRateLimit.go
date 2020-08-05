package main

import (
	"fmt"
	"time"
)

func main() {
	reqs := make(chan int, 5)
	go func() {
		for i := 0; i < 50; i++ {
			reqs <- i
		}
		// close之后more才会为false，否则一直等待
		close(reqs)
	}()
	limiter := time.Tick(time.Millisecond * 100)
	result := make(chan bool)
	go func() {
		for {
			// 每100ms取一次，通过Tick实现
			t := <-limiter
			req, more := <-reqs
			if more {
				fmt.Println(t, req)
			} else {
				fmt.Println("done")
				result <- true
				return
			}
		}
	}()
	<-result
}
