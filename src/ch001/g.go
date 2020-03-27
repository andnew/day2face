package ch001

import "fmt"

func MainCall1() {
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

// 总结&分析
// 主要考察的知识点有
// defer 和 panic
// defer 表示 延迟执行 ; panic 表示 异常 中断

// 本面试题的执行结果是
// 打印后
// 打印中
// 打印前
// panic: 触发异常
