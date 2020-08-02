package main

import "fmt"

func main() {
	i := 1
	zero(i)
	fmt.Println(i)
	zero2(&i)
	fmt.Println(i)
	fmt.Println(&i)
}

func zero(i int) {
	i = 0
}

func zero2(i *int) {
	*i = 0
}
