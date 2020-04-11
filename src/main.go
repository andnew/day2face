package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover()) // 对应的是 panic(1)
	}()
	defer func() {
		defer fmt.Print(recover()) // 对应的是 panic(2)
		panic(1)
	}()
	defer recover()
	panic(2)
}

// 执行结果 21

//func main() {
//	defer func() {
//		fmt.Print(recover())
//	}()
//	defer func() {
//		defer func() {
//			fmt.Print(recover())
//		}()
//		panic(1)
//	}()
//	panic(2)
//}
// 执行结果 12

// 总结&分析
// 知识点：defer recover ,recover 必须放在defer 函数内
// recover() 必须在 defer() 函数中调用才有效，所以第 9 行代码捕获是无效的。在调用 defer() 时，便会计算函数的参数并压入栈中，所以执行第 6 行代码时，此时便会捕获 panic(2)；此后的 panic(1)，会被上一层的 recover() 捕获。所以输出 21。
