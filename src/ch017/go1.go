package ch017

import "fmt"

func MainCall1() {

	// 示例 1
	demo11()
	// 示例 2
	demo12()
	// 示例3
	demo13()
}

func demo11() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Printf("a length is %d, cap is %d\n", len(a), cap(a))
	fmt.Printf("b length is %d, cap is %d\n", len(b), cap(b))
	fmt.Printf("c length is %d, cap is %d\n", len(c), cap(c))
}

// 总结&分析
// 切片的长度和容量的计算问题
// 数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。
// 在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
// 操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。

// 下面代码中 A B 两处应该怎么修改才能顺利编译？
func demo12() {
	//var m map[string]int        //A
	var m map[string]int = make(map[string]int)
	m["a"] = 1
	//if v := m["b"]; v != nil {  //B
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}

// 总结&分析
// 编译报错 cannot convert nil to type int
// 主要考察的是 map 的声明、赋值 和 数据的获取

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func demo13() {
	c := Work{3}
	var a A = c
	var b B = c
	//fmt.Println(a.ShowB())
	//fmt.Println(b.ShowA())
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

// 总结&分析
// 考察的是接口、结构体和多态
