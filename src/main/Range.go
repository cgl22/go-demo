package main

import "fmt"

func main() {
	// slice循环
	num := []int{1, 2, 3, 4}
	sum := 0
	for _, n := range num {
		sum += n
	}
	fmt.Println(sum)

	for i, _ := range num {
		fmt.Println("index", i)
	}

	// map循环s
	m := map[string]int{"k1": 1, "k2": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, _ := range m {
		fmt.Println(k)
	}

	for i, s := range "hello go" {
		fmt.Println(i, s)
	}
}
