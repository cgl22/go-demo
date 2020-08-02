package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, i := range []int{1, 3, 4} {
		if result, er := error1(i); er == nil {
			fmt.Println(result)
		} else {
			fmt.Println(er.Error())
		}
	}

	for _, i := range []int{1, 3, 4} {
		if result, er := error2(i); er == nil {
			fmt.Println(result)
		} else {
			fmt.Println(er.Error())
		}
	}

	// ???
	_, e := error2(1)
	fmt.Println(e)
	if ae, ok := e.(*myError); ok {
		fmt.Println(ae)
	}
}

// 支持自定义错误
type myError struct {
	code    int
	message string
}

// 需要实现Error接口
func (er *myError) Error() string {
	return "自定义error实现：" + er.message
}

func error2(a int) (int, error) {
	if a == 1 {
		return -1, &myError{-1, "1真的不行"}
	}
	return a + 1, nil
}

// 按照惯例，错误通常是最后一个返回值并且是error类型(一个内建的接口)
func error1(a int) (int, error) {
	if a == 1 {
		return a, errors.New("1不行啊")
	}
	// nil代表无错误
	return a + 1, nil
}
