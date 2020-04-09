package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v", r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	f()
}

// 总结&分析
// 知识点：panic、recover()。当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容
