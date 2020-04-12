package main

import "fmt"

//type A interface {
//	ShowA() int
//}
//
//type B interface {
//	ShowB() int
//}
//
//type Work struct {
//	i int
//}
//
//func (w Work) ShowA() int {
//	return w.i + 10
//}
//
//func (w Work) ShowB() int {
//	return w.i + 20
//}
//
//func main1() {
//	c := Work{3}
//	var a A = c
//	var b B = c
//	fmt.Println(a.ShowA())
//	fmt.Println(b.ShowB())
//}
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

func main2() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
