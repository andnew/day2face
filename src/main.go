package main

import "fmt"

func change(s ...int) {
	fmt.Println("len = ", len(s), " , cap ", cap(s), " s = ", s)
	s = append(s, 3)
	fmt.Println("len = ", len(s), " , cap ", cap(s), " s = ", s)
}

func main1() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

// 总结&分析
// 执行结果
// len =  5  , cap  5  s =  [1 2 0 0 0]
// len =  6  , cap  10  s =  [1 2 0 0 0 3]
// [1 2 0 0 0]
// len =  2  , cap  5  s =  [1 2]
// len =  3  , cap  5  s =  [1 2 3]
// [1 2 3 0 0]

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12 // 1，12，3, 4, 5
			a[2] = 13 // 1，12，13, 4, 5
		}
		r[i] = v // 1, 12, 13, 4, 5
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
