package main

import "fmt"

var x = []int{2: 2, 3, 0: 1}

func main1() {
	fmt.Println(x)
}

// 执行结果 [1 0 2 3]

func incr(p *int) int {
	*p++
	return *p
}
func main() {
	v := 1
	incr(&v)
	fmt.Println(v) // 2
}
