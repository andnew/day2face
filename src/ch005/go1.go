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
