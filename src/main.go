package main

import "fmt"

type X struct{}

func (x *X) test() {
	println(x)
}

func main1() {
	var a *X
	a.test()

	(&X{}).test()
}
func main2() {
	var a *X
	a.test() // 相当于 test(nil)

	var x = X{}
	x.test()
}

func main() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	//if v := x["two"]; v == "" {
	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	}
}
