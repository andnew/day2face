package main

import "fmt"

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

func main1() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

func main2() {
	s := [3]int{1, 2, 3}
	a := s[:0]         // 0 3
	b := s[:2]         // 2 3
	c := s[1:2:cap(s)] // 1 2
	fmt.Println("a len ", len(a), ", cap ", cap(a))
	fmt.Println("b len ", len(b), ", cap ", cap(b))
	fmt.Println("c len ", len(c), ", cap ", cap(c))
}

func main() {
	var m map[string]int = make(map[string]int) //A
	m["a"] = 1
	if v, ok := m["b"]; ok { //B
		fmt.Println(v)
	}
}
