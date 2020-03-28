package ch003

import "fmt"

func MainCall2() {
	s := make([]int, 0)
	fmt.Printf("append 前，类型: %T , 长度: %d , 容量: %d \n", s, len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("append 后，类型: %T , 长度: %d , 容量: %d \n", s, len(s), cap(s))
	fmt.Println(s)
}

//总结&分析
// 同样是考察 切片的，唯一的区别 初始化 长度是 0
