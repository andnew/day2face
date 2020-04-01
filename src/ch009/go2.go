package ch009

import "fmt"

func MainCall2() {
	demo21()
}

func printWord(ch chan string) {
	fmt.Println("Hello " + <-ch)
}

func productCh(ch chan chan string) {
	c := make(chan string) // 创建 string type 信道
	ch <- c                // 传输信道
}

func demo21() {
	// 创建 chan string 类型的信道
	ch := make(chan chan string)
	go productCh(ch)
	// c 是 string type 的信道
	c := <-ch
	go printWord(c) //这是ch信道处于堵塞状态
	c <- "world"    //c 信道写入 ，ch 信道 堵塞 解除，从 ch 读出数据
	fmt.Println("main stopped")
}

// 总结&分析
// 使用通道作为另一个通道的数据类型
