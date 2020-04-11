package main

import "fmt"

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}

type Integer int

//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

// 总结&分析
// 值接收者的方法可以使用值或者指针调用
// 对于指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
