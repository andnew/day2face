package main

import "fmt"

func main1() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	_ = y == y // 发生panic
}

// 总结&分析
// 因为两个比较值的动态类型为同一个不可比较类型。

var o = fmt.Print

func main() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- 1:
			o(3)
		}
	}
}

// 总结&分析
// 第一次循环，写操作已经准备好，执行 o(3)，输出 3；
// 第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；
// 第三次，由于 c 为 nil，走的是 default 分支，输出 1。
// 执行结果是 321
