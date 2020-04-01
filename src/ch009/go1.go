package ch009

import "fmt"

func MainCall1() {
	// 示例1
	demo11()
	// 示例2
	demo12()
	// 示例3
	demo13()
	// 示例4
	demo14()

}

func hello(num ...int) {
	num[0] = 18
}
func demo11() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

type person struct {
	name string
}

func demo12() {
	var m map[person]int
	p := person{"mike"}

	fmt.Println(m == nil, m[p])
}

func change(s ...int) {
	s = append(s, 3)
	fmt.Println("function change ", s)
}

func demo13() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println("第一次打印 change 后打印 ", slice)
	change(slice[0:2]...)
	fmt.Println("第二次打印 change 后打印 ", slice)
	// 执行结果
	//function change  [1 2 0 0 0 3]
	//第一次打印 change 后打印  [1 2 0 0 0]
	//function change  [1 2 3]
	//第二次打印 change 后打印  [1 2 3 0 0]

}

// 总结&分析
// 第一次change 传递的参数,是 slice 整个切片, change 函数 append 3 ，导致 函数内的 s 执行的 新的切片地址
// 第二次change 传递的参数,是 slice [0:2}切片, change 函数 append 3 ,作用在原来的切片上

func change1(s ...string) {
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
func demo14() {
	slice1 := []string{"Hello", "World", "Go"}
	slice2 := []string{"Aa", "Bb"}
	// 1、
	change1(append(slice1, "Again")...)
	// 2、
	change1(append(slice1, slice2...)...)

	//执行结果
	//[]string
	//[Hello World Go Again]
	//[]string
	//[Hello World Go Aa Bb]
}

// 总结&分析
// 1、append 第二个或之后的是可变参数
// 2、可变函数，作为入参时，如 change1(s ...string),类型前必须有 3 个 .
// 3、可变函数, 调用函数，函数的参数是可变函数，调用时 如果参数 是 切片的话，作为入参，必须在参数加入 ...
