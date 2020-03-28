package ch004

import "fmt"

func MainCall2() {
	s1 := []int{1, 2, 3}
	fmt.Printf("append 前, s1 的数据类型:%T, 长度:%d, 容量:%d\n", s1, len(s1), cap(s1))
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Printf("append 后, s1 的数据类型:%T, 长度:%d, 容量:%d\n", s1, len(s1), cap(s1))
	fmt.Println(s1)
}

// 总结&分析
// 编译报错 cannot use s2 (type []int) as type int in append
// 变量 s1 和 s2 都是切片， append 操作的第二个或后面的数据，必须是元素，如果是切片的话，必须使用 ... 处理
