package main

var f = func(i int) {
	print("x")
}

func main() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}

// 总结&分析
// 闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。
