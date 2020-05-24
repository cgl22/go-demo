package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	output := "0~100sum为:"
	fmt.Println(output + strconv.Itoa(sum))

	// 类似while
	sum = 0
	for sum <= 10 {
		sum++
	}
	fmt.Println(sum)

	// 类似while true
	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)

	zss := int64(99)
	var zsd int64 = 10
	fmt.Println(zss + zsd)

	// foreach
	numbers := []int32{1, 5, 1, 8}
	for i, j := range numbers {
		fmt.Printf("index为%d，值为%d\n", i, j)
	}
	for j := range numbers {
		fmt.Printf("值为%d\n", j)
	}
}
