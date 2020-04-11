package main

import "fmt"

const i = 100

var j = 123

func main1() {
	fmt.Println(&j, j)
	//fmt.Println(&i, i) //cannot take the address of i
}

func GetValue(m map[int]string, id int) (string, bool) {

	if _, exist := m[id]; exist {
		return "exist", true
	}
	return "", false // 正确的代码 return "" , false
}
func main() {
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}
