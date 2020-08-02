package main

import "fmt"

func main() {
	a := make(map[string]string)

	a["key1"] = "value1"
	a["key2"] = "value2"
	str := "key1"
	fmt.Println(a)
	fmt.Println(a[str])
	fmt.Println(len(a))

	// 删除
	delete(a, str)
	fmt.Println(a)

	// 分别输出“”，false
	c := a[str]
	fmt.Println(c)
	_, b := a[str]
	fmt.Println(b)

	// 声明并初始化
	d := map[string]string{"k1": "v1", "k2": "v2"}
	fmt.Println(d)
}
