package main

import (
	"fmt"
	"runtime/debug"
)

func f(n int) (r int) {
	defer func() {
		r += n
		if err := recover(); err != nil {
			fmt.Println("9999")
			debug.PrintStack()
		}
	}()

	var f func()

	//defer f() // 未定义，引发异常
	f = func() {
		fmt.Println("前======", r)
		r += 2
		fmt.Println("后======", r)
	}
	defer f()
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
