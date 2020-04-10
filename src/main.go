package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("---%#v", o.Quantity)
}

func main() {
	//s := NewSlice()
	//defer func() {
	//	s.Add(1).Add(2)
	//}()
	//s.Add(3)

	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}

// 总结&分析
// 闭包延迟求值。for 循环局部变量 i，匿名函数每一次使用的都是同一个变量。
