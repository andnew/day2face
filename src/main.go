package main

import "fmt"

//type T struct {
//	ls []int
//}
//
//func foo(t T) {
//	t.ls[0] = 100
//}
//func main() {
//	var t = T{
//		ls: []int{1, 2, 3},
//	}
//	foo(t)
//	fmt.Println(t.ls[0])
//}

func main() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}

// 总结&分析
// 执行结果
//false
//true
