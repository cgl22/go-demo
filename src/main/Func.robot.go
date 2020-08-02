package main

import "fmt"

type cb func(int) int

func main() {
	a := 10
	b := 20
	var m int = max(a, b)
	fmt.Println(m)

	x, y := mutilReturn("first", "second")
	fmt.Println(x, y)

	swap(&a, &b)
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(a, b))

	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	// 可变参数
	fmt.Println(sum(1, 2, 3))
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 返回多个值
func mutilReturn(a, b string) (string, string) {
	return b, a
}

// 传地址引用，默认为传值引用
func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
