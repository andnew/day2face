package main

import "fmt"

//type T struct {
//	n int
//}
//
//func main() {
//	ts := [2]T{}
//	for i, t := range ts  {
//		switch i {
//		case 0:
//			t.n = 3
//			ts[1].n = 9
//		case 1:
//			fmt.Print(t.n, " ")
//		}
//	}
//	fmt.Print(ts)
//}

// 总结&分析
// for-range 循环数组。此时使用的是数组 ts 的副本，所以 t.n = 3 的赋值操作不会影响原数组。
// 执行结果
// 0 [{0} {9}]

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

// 总结&分析
// 知识点：for-range 数组指针。for-range 循环中的循环变量 t 是原数组元素的副本。如果数组元素是结构体值，则副本的字段和原数组字段是两个不同的值。
// 参考答案及解析：9 [{0} {9}]。
