package main

import "fmt"

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main1() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D
}

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age) // 28

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age) // 28
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age) // 29
	}()

	person = &Person{29}
}
