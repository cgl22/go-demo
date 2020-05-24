package main

import "fmt"

func main() {
	// if
	var age int = 18
	if age == 18 {
		fmt.Println("age为18")
	} else {
		fmt.Println("age不为18")
	}

	// switch
	var score int = 80
	switch score {
	case 60:
		fmt.Println("及格")
	case 70:
		fmt.Println("一般")
	case 80:
		fmt.Println("良好")
	case 90:
		fmt.Println("优秀")
	default:
		fmt.Println("确定不了")
	}
	switch score == 80 {
	case true:
		fmt.Println("80分")
	case false:
		fmt.Println("不是80分")
	default:
		fmt.Println("不知道多少分")
	}
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score < 70:
		fmt.Println("及格")
	case score < 80:
		fmt.Println("一般")
	case score < 90:
		fmt.Println("良好")
		break
		fallthrough
	case score == 10000:
		fmt.Println("fallthrough 跳转到这里了")
	case score < 100:
		fmt.Println("优秀")
	default:
		fmt.Println("非常人")
	}

	// select
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}
