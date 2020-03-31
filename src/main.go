package main

import (
	ch "ch008"
	"fmt"
)

func init() {
	fmt.Println("package main function init ")
}

func init() {
	fmt.Println("package main function init 2")
}

func main() {
	// 试题1
	ch.MainCall1()
	// 常量示例
	ch.MainCall2()
	// 包级别执行
	ch.MainCall3()

}
