package main

import "fmt"

func main1() {
	//list := new([]int)
	// 正确的代码
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
}

// 编译报错 Cannot use 'list' (type *[]int) as type []Type

func main2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

// 编译报错  Cannot use 's2' (type []int) as type int

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}
