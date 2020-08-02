package main

import "fmt"

type student struct {
	name string
	age  int
}

// struct的方法
func (s student) doSth(param int) {
	fmt.Println("student的方法")
}

func main() {
	fmt.Println(student{"张三", 9})
	fmt.Println(student{age: 6})
	fmt.Println(&student{name: "李四", age: 10})

	s := student{"王五", 9}
	fmt.Println(s.age)

	// 等效,指针会被自动解引用
	ss := &s
	fmt.Println(ss.name)
	sss := s
	fmt.Println(sss.name)

	sss.name = "hahahhaha"
	fmt.Println(sss.name)

	sss.doSth(99)
}
