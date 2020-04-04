package ch104

import "fmt"

func MainCall1() {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}

func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}

// 总结&分析
// 1、defer 使用，执行顺序 return 执行之后，返回值 之前
// 2、函数的返回值参数的 声明 化
// 3、函数的 return value 不是原子操作，而是在编译器中分解为两部分：返回值赋值 和 return
// 执行结果
// 0
// 40
// 50
