package main

import "fmt"

//type N int
//
//func (n N) value() {
//	n++
//	fmt.Printf("v:%p,%v\n", &n, n)
//}
//func (n *N) pointer() {
//	*n++
//	fmt.Printf("v:%p,%v\n", n, *n)
//
//}
//func main() {
//	var a N = 25
//	p := &a
//	p1 := &p
//	(*p1).value()
//	(*p1).pointer()
//}

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	fmt.Println(n) // 10

	n++ //11
	f1 := N.test
	f1(n) // 11

	n++ // 12
	f2 := (*N).test
	f2(&n) //12
}
