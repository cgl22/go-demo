package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int, 10)
	tasks := make(chan int, 10)

	// 启动三个worker
	for i := 0; i < 3; i++ {
		go worker(tasks, result, i)
	}
	// 生成9个任务
	for i := 0; i < 9; i++ {
		tasks <- i
	}
	close(tasks)
	for i := 0; i < 9; i++ {
		fmt.Println(<-result)
	}
	// 由于有3个worker同时执行任务，执行时间小于9s，接近3s
}

// worker从tasks接收任务，并将结果返回给result
// 每个任务消耗1s
func worker(tasks <-chan int, result chan<- int, id int) {
	for t := range tasks {
		time.Sleep(time.Second)
		fmt.Println("worker_id:", id, " task:", t)
		result <- t
	}
}
