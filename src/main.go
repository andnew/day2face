package main

import "fmt"

//type T struct {
//	n int
//}
//
//func main() {
//	ts := [2]T{}
//	for i := range ts[:] {
//		switch i {
//		case 0:
//			ts[1].n = 9
//		case 1:
//			fmt.Print(ts[i].n, " ")
//		}
//	}
//	fmt.Print(ts)
//}

// 总结&分析
// 切片 指针
// for-range 切片时使用的是切片的副本，但不会复制底层数组，换句话说，此副本切片与原数组共享底层数组。

type T struct {
	n int
}

func main() {
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
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
// for 切片 指针
// 9 [{3} {9}]
