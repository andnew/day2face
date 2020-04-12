package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

// 总结
// defer 先进后出 ，panic 的错误，需要 recover ,但是recover 必须放在defer 函数中
// 执行结果如下
// 打印后
// 打印中
// 打印前
// panic: 触发异常
