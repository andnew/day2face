package main

import "fmt"

type N int

func (n *N) test() {
	fmt.Println(*n)
}

func main1() {
	var n N = 10
	p := &n

	n++ // 11
	f1 := n.test

	n++ //12
	f2 := p.test

	n++            //13
	fmt.Println(n) //13

	f1() //13
	f2() //13
}

// 总结&分析
// 方法值。当目标方法的接收者是指针类型时，那么被复制的就是指针。

func main() {
	var m map[int]bool // nil
	_ = m[123]
	var p *[5]string // nil
	for range p {
		_ = len(p)
	}
	var s []int // nil
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
}
