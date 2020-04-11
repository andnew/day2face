package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x) // 1
	{
		fmt.Println(x)    // 1
		i, x := 2, 2      // i = 2 , x = 2
		fmt.Println(i, x) // 2 , 2
	}
	fmt.Println(x) // print ?  1
}

// 总结&分析
// 变量隐藏。使用变量简短声明符号 := 时，如果符号左边有多个变量，只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。
// 但如果出现作用域之后，就会导致变量隐藏的问题
