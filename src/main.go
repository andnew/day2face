package main

import "fmt"

type Foo struct {
	bar string
}

func main1() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		//fmt.Printf("%p\n", &value)
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])
}

// 总结&分析
// 值接收者的方法可以使用值或者指针调用
// 对于指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
// 修复代码
// for i := range s1 {
//    s2[i] = &s1[i]
// }

func main() {

	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	for {
		counter := 0
		for k, v := range m {
			if counter == 0 {
				delete(m, "A")
			}
			counter++
			fmt.Println(k, v)
		}
		fmt.Println("counter is ", counter, ", m is ", m)

		if counter == 2 {
			break
		}

	}
}
