package main

import "fmt"

//func main() {
//	var k = 1
//	var s = []int{1, 2}
//	k, s[k] = 0, 3
//	fmt.Println(s[0] + s[1])
//}

// 总结&分析
// k, s[k] = 0, 3 ==> k, s[1] = 0, 3

func main() {
	var k = 9
	for k = range []int{} {
		fmt.Println("=====")
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
		fmt.Println("----")
	}
	fmt.Println(k)

	for k = range (*[3]int)(nil) {
		fmt.Println("666--")
	}
	fmt.Println(k)
}

// 执行结果
//9
//3
//2
