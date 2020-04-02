package ch107

import (
	"fmt"
	"runtime"
)

func MainCall1() {
	demo11()
}

// 一段耗时的计算函数
func consumer(ch chan int) {

	// 无限获取数据的循环
	for {

		fmt.Println("channel----------", &ch)
		// 从通道获取数据
		data := <-ch //没有数据向通道写入，导致读操作一直处于堵塞状态

		// 打印数据
		fmt.Println("========== ", data)
	}

}

func demo11() {

	// 创建一个传递数据用的通道
	ch := make(chan int)

	for {

		// 空变量, 什么也不做
		var dummy string

		// 获取输入, 模拟进程持续运行
		fmt.Scan(&dummy)

		// 启动并发执行consumer()函数
		go consumer(ch)

		// 输出现在的goroutine数量
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}

// 总结&分析
// 这个程序实际在模拟一个进程根据需要创建 goroutine 的情况。
// 运行后，问题已经被暴露出来：随着输入的字符串越来越多，goroutine 将会无限制地被创建，但并不会结束。
// 这种情况如果发生在生产环境中，将会造成内存大量分配，最终使进程崩溃。
