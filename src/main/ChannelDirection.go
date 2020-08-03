package main

import "fmt"

func main() {
	producer := make(chan string, 1)
	consumer := make(chan string, 1)
	produce(producer, "msg")
	consume(producer, consumer)
	fmt.Println(<-consumer)
}

// producer只能发送数据
func produce(producer chan<- string, msg string) {
	producer <- msg
}

// consumer只能接收数据，producer只能发送数据
func consume(consumer <-chan string, producer chan<- string) {
	msg := <-consumer
	producer <- msg
}
