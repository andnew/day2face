package main

import "fmt"

type T struct {
	x int
	y *int
}

func main1() {
	i := 20
	t := T{10, &i}
	p := &t.x
	*p++
	*p--
	t.y = p
	fmt.Println(*t.y)
}

func main() {
	x := make([]int, 2, 10)
	_ = x[6:10]
	_ = x[6:]
	_ = x[2:]
}

// 总结&分析
// 截取符号 [i:j]，如果 j 省略，默认是原切片或者数组的长度，x 的长度是 2，小于起始下标 6 ，所以 panic。
