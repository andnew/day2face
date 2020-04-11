package main

import "fmt"

var gvar int

func main1() {
	var one int
	two := 2
	var three int
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")

	fmt.Println(one, two, three)
}

// 变量声明与使用

//type ConfigOne struct {
//	Daemon string
//}
//
//func (c *ConfigOne) String() string {
//	return fmt.Sprintf("print: %v", c)
//}
//
//func main() {
//	c := &ConfigOne{}
//	c.String()
//}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println(r, a)
}
