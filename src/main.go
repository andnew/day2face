package main

import "fmt"

type person struct {
	name string
}

func main1() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

func hello(num ...int) {
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}
	//hello(i...)
	hello(i[:]...)
	fmt.Println(i[0])
}
