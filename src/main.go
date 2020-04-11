package main

import "fmt"

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func main() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)
}

// 总结&分析
// 执行结果 7 6 8
// 知识点，defer 执行顺序 ，先进后出， 值传递 和 匿名函数
