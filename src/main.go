package main

import "fmt"

func main1() {
	s := make([]int, 3, 9)
	fmt.Println(len(s)) //3
	s2 := s[4:8]
	fmt.Println(len(s2)) //4
}

// 总结&分析
// 执行结果是 3 4

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	p := &n
	n++
	f1 := n.test
	n++
	f2 := p.test
	n++
	fmt.Println(n)
	f1()
	f2()
}
