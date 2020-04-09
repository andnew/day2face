package main

import "fmt"

func main() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		//fmt.Println("m = ", m, " , n = ", n, " , i = ", i)
		if n == -n {
			fmt.Println("n-----------", n, "---------------")
			count++
		}
		if m == -m {
			fmt.Println("m-----------", m, "---------------")
			count++
		}
	}
	fmt.Println(count) // 4
}

// 总结&分析
// 知识点：数值溢出。当 i 的值为 0、128 是会发生相等情况，注意 byte 是 uint8 的别名。
