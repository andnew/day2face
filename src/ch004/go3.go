package ch004

import "fmt"

var (
	size     = 1024
	max_size = size * 2
)

func MainCall3() {
	fmt.Println(size, max_size)
}

// 总结&分析
// 编译错误 syntax error: unexpected :=, expecting =
// 变量初始化的2个方式 ，2个操作，声明和赋值
// var 是 定义变量的关键字 ； 其中 := 采用 类型推导的方式定义 变量
