package main

import "fmt"

type animal interface {
	sound()
	walk()
}

type duck struct {
	weight int
}

type dog struct {
	legs int
}

func (d duck) sound() {
	fmt.Println(d.weight, "嘎嘎嘎")
}

func (d duck) walk() {
	fmt.Println("wwwww")
}

func interfaceTest(a animal) {
	a.sound()
	a.walk()
}

func main() {
	duck := duck{99}
	interfaceTest(duck)
}
