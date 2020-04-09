package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		j := i
		funs = append(funs, func() {
			println(&j, j)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

// 总结&分析
// 闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。
