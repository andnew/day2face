package ch008

import (
	"fmt"
)

var precomputed = [20]float64{}

func init() {
	fmt.Println("package ch008 function init of go1.go ...")
	var current float64 = 1
	precomputed[0] = current
	for i := 1; i < len(precomputed); i++ {
		precomputed[i] = precomputed[i-1] * 1.2
	}
}
func init() {
	fmt.Println("package ch008 function init of go1.go ..... 2 ")
}

func MainCall1() {
	// 试题1
	//demo01()
	// 试题2
	demo02()
	fmt.Println(precomputed)
}

func GetValue() int {
	return 1
}

//func demo01() {
//	i := GetValue()
//	//i.(type)
//	switch reflect.TypeOf(i) {
//	case int:
//		println("int")
//	case string:
//		println("string")
//	case interface{}:
//		println("interface")
//	default:
//		println("unknown")
//	}
//}
// 总结&分析
// 类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择

func hello() []string {
	return nil
}
func demo02() {
	//h := hello
	h := hello()

	if h == nil {
		fmt.Println("nil") // h := hello() //获取的是函数的返回值
	} else {
		fmt.Println("not nil") // h := hello // h是函数
	}
}

// 总结&分析
// hello 代表的是一个函数 ,返回 func []string

//关于init函数，下面说法正确的是(AB)
//A. 一个包中，可以包含多个 init 函数；
//B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
//C. main 包中，不能有 init 函数；
//D. init 函数可以被其他函数调用；
