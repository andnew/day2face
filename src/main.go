package main

import "fmt"

//func main() {
//	var fn1 = func() {}
//	var fn2 = func() {}
//
//	if fn1 != fn2 {
//		println("fn1 not equal fn2")
//	}
//}
// Invalid operation: fn1 != fn2 (operator != not defined on func())

type T struct {
	n int
}

func main() {
	m := make(map[int]T)
	t := T{1}
	m[0] = t
	fmt.Println(m[0].n)
}

// 总结&分析
// 编译错误：Cannot assign to m[0].n
// map[key]struct 中 struct 是不可寻址的，所以无法直接赋值。
//
