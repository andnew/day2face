package ch005

import "fmt"

func MainCall1() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1.age == sm2.age {
		fmt.Println("sm1.age == sm2.age")
	}
}

// 总结&分析
// 编译错误  invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
// map 不可以使用 == 进行比较 , 结构体的元素可以 == 一一比较
// 如果结构体包含不可比较的字段，则结构体变量也不可比较。
// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
// 判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素

// 官方总结
//参考答案及解析：编译不通过 invalid operation: sm1 == sm2
//
//这道题目考的是结构体的比较，有几个需要注意的地方：
//
//结构体只能比较是否相等，但是不能比较大小。
//
//相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体；
//
//1    sn3:= struct {
//2        name string
//3        age  int
//4    }{age:11,name:"qq"}
//如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
//
//那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。 具体可以参考 Go 说明文档。https://golang.org/ref/spec#Comparison_operators
