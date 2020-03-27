package ch003

import "fmt"

func MainCall1() {
	s := make([]int, 5)
	fmt.Printf("append 前，类型: %T , 长度: %d , 容量: %d \n", s, len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("append 后，类型: %T , 长度: %d , 容量: %d \n", s, len(s), cap(s))
	fmt.Println(s)
}

// 总结&分析
// 其中 append 前/后 信息的打印，是我们自己添加进去的
// 本题主要的考察的知识点是 切片
