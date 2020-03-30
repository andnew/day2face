package ch007

import "fmt"

//试题
//1.关于字符串连接，下面语法正确的是？ B
//A. str := 'abc' + '123'
//B. str := "abc" + "123"
//C. str := '123' + "abc"
//D. fmt.Sprintf("abc%d", 123)

func MainCall1() {
	// 示例01
	demo01()
	// 示例02
	demo02()
	// 示例03
	demo03()
}

func demo01() {
	abc := fmt.Sprintf("abc%d", 123)
	fmt.Println(abc)
}

// 总结&分析
// 字符串定义

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func demo02() {
	fmt.Println(x, y, z, k, p)
}

// 总结&分析
// 常量的定义

//3.下面赋值正确的是(BD)
//A. var x = nil
//B. var x interface{} = nil
//C. var x string = nil
//D. var x error = nil

func demo03() {
	var x error = nil

	fmt.Println(x)
}
