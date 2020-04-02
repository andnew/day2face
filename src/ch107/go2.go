package ch107

import (
	"fmt"
	"runtime"
)

func MainCall2() {
	demo21()
}

// 针对go1的例子，进行优化
// 设置了一个退出的条件，但是条件永远不会被满足或者触发。

// 一段耗时的计算函数
func consumer2(ch chan int) {
	// 无限获取数据的循环
	for {
		// 从通道获取数据
		fmt.Println("---------等待读取数据channel-------", &ch)
		data := <-ch
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println("go2--------", data, &ch)
	}
	fmt.Println("go2--------goroutine exit", &ch)
}
func demo21() {
	// 传递数据用的通道
	ch := make(chan int)
	for {
		// 空变量, 什么也不做
		var dummy string
		// 获取输入, 模拟进程持续运行
		fmt.Scan(&dummy)
		if dummy == "quit" {
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				ch <- 0
			}
			continue
		}
		// 启动并发执行consumer()函数
		go consumer2(ch)
		// 输出现在的goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
