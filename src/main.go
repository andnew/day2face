package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// 总结&分析
// for-range 切片 ，其中for key value 是 内部定义的变量，保持不变
// 执行结果
// 0 -> 3
// 1 -> 3
// 2 -> 3
// 3 -> 3

// 修正代码
// 	for key, val := range slice {
//      var value = val
//		m[key] = &value
//	}
