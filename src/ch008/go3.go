package ch008

import "fmt"

var (
	a int = b + 1
	b int = 1
)

//var (
//	a1 = b1
//	b1 = c1
//	c1 = f()
//)
//func f() int {
//	return a1
//}

func MainCall2() {
	//
	demo31()
	//
	//demo32()

	//fmt.Println(a1, b1, c1) // initialization loop:
}

func demo31() {
	fmt.Println(a) // 2
	fmt.Println(b) // 1
}

//func demo32() {
//	var (
//		c int = d + 1
//		d int = 1
//	)
//	fmt.Println("c=", c)
//	fmt.Println("d=", d) // 编译 undefined: d
//}
