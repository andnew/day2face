package ch03

import "fmt"

//方式1 提供一个参数
func funMui(x, y int) (sum int, err error) {
	//return x + y, nil
	sum = x + y
	err = nil
	return sum, err
}

// 方式2 省去参数名
func funMui1(x, y int) (int, error) {
	return x + y, nil
}

func MainCall3() {
	sum, ok := funMui(10, 20)
	if ok != nil {
		fmt.Println(sum)
	} else {
		fmt.Println("funMui计算的结果是:", sum)
	}

}

// 总结&分析
// 编译报错 syntax error: mixed named and unnamed function parameters
// 报错的原因是 func funMui(x, y int) (sum int, error) error 没有参数名导致
// 针对这种混合的参数处理方式如下
// 方式1、func funMui(x, y int) (sum int,err error)  给它一个参数名
// 方式2、func funMui(x, y int) (int, error) 省去参数
// 其中方式1的函数体存在多种处理方式
// 处理方式1、
//   func funMui(x, y int) (sum int,err error) {
//	    return x + y, nil
//   }
// 处理方式2、
//   func funMui(x, y int) (sum int,err error) {
//	    sum = x + y
//	    err = nil
//      return sum, nil // return 必不可少的
//   }
