package ch006

import "fmt"

// 1.通过指针变量 p 访问其成员变量 name，有哪几种方式？
// A.p.name
// B.(&p).name
// C.(*p).name
// D.p->name
// 正确答案是 AC
// 总结&分析
// & 取址运算符，* 指针解引用。

func MainCall1() {
	var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

type MyInt1 int
type MyInt2 = int

// 总结&分析
// 报错 cannot use i (type int) as type MyInt1 in assignment
