package chapter

import "fmt"

func MainCall1() {
	demo11()
}

type data struct {
	name string
}

func (p data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func demo11() {
	d1 := data{"one"}
	d1.print()
	var in printer = data{"two"}
	in.print()
}

// 总结&分析
// 结构体多态的 主要考虑到接口的实现 使用的对象 T
// 结构体添加方法 作用在对象的指针上 *T