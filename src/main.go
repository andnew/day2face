package main

import (
	"fmt"
	"os"
	"strings"
)

func main1() {
	var a []int = make([]int, 1)
	a, a[0] = []int{1, 2}, 9
	fmt.Println(a)
}

const (
	azero = iota
	aone  = iota
)

const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func main2() {
	fmt.Println(azero, aone) // 0 1
	fmt.Println(bzero, bone) // 1 2
}

//下面这段代码输出什么？

func main3() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			fmt.Println("n== ", n)
			count++
		}
		if m == -m {
			fmt.Println("m== ", m)
			count++
		}
	}
	fmt.Println(count)
}

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main4() {
	d1 := data{"one"}
	d1.print()

	var in printer = &data{"two"}
	in.print()
}

func main5() {
	a := [3]int{0, 1, 2}
	s := a[1:2] // {1} 1 , 2

	s[0] = 11         // {11}
	s = append(s, 12) //s = {11,12} , a = {0,11,12}
	s = append(s, 13) //s = {11,12,13} , a = {0,11,12}
	s[0] = 21         //{21,12,13}

	fmt.Println(a) //{0,11,12}
	fmt.Println(s) //{21,12,13}
}

func main6() {
	fmt.Println(strings.TrimRight("ABBA", "B"))
}

func main7() {
	n := 43210
	fmt.Println(n/(60*60), " hours and ", n%(60*60), " seconds")
}

const (
	Century = 100
	Decade  = 010
	Year    = 001
)

func main8() {
	fmt.Println(Century + 2*Decade + 2*Year)
}

func main9() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}

func Foo() error {
	var err *os.PathError = nil
	// …
	return err
}

func main10() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

type T []int

func F(t T) {
	fmt.Println(t)
}

func main11() {
	var q []int = []int{1, 2, 3}
	F(q)
}

func min(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func main12() {
	min(1225, 256)
}

// 下面哪些函数不能通过编译？

func A(string string) string {
	return string + string
}

func B(len int) int {
	return len + len
}

func C(val, d string) string {
	if val == "" {
		return d
	}
	return val
}

// 下面代码输出什么？请简要说明。
type foo struct{ Val int }
type bar struct{ Val int }

func main() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Print(a == b, c == foo(d), e == f)
}
