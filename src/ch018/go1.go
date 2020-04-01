package ch018

import "fmt"

func MainCall1() {
	fmt.Println("f1()的返回值:", f1())
	fmt.Println("f2()的返回值:", f2())
	fmt.Println("f3()的返回值:", f3())
	// 执行结果
	// f1()的返回值: 1
	// f2()的返回值: 5
	// f3()的返回值: 1

}

func f1() (r int) {
	fmt.Printf("f1()=== func 外部前,%v , 地址 %p\n", r, &r)
	defer func() {
		fmt.Printf("f1()=== func 内部前,%v , 地址 %p\n", r, &r)
		r++
		fmt.Printf("f1()=== func 内部后,%v , 地址 %p\n", r, &r)
	}()
	fmt.Printf("f1()=== func 外部后,%v , 地址 %p\n", r, &r)
	return 0
	// 方法内部的执行
	//f1()=== func 外部前,0 , 地址 0xc00001c090
	//f1()=== func 外部后,0 , 地址 0xc00001c090
	//f1()=== func 内部前,0 , 地址 0xc00001c090
	//f1()=== func 内部后,1 , 地址 0xc00001c090
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	fmt.Printf("f3()=== func 外部前,%v , 地址 %p\n", r, &r)
	defer func(r int) {
		fmt.Printf("f3()=== func 内部前,%v , 地址 %p\n", r, &r)
		r = r + 5
		fmt.Printf("f3()=== func 内部后,%v , 地址 %p\n", r, &r)
	}(r)
	fmt.Printf("f3()=== func 外部后,%v , 地址 %p\n", r, &r)
	return 1

	// 方法内部的执行
	//f3()=== func 外部前,0 , 地址 0xc00001c0b0
	//f3()=== func 外部后,0 , 地址 0xc00001c0b0
	//f3()=== func 内部前,0 , 地址 0xc00001c0b8
	//f3()=== func 内部后,5 , 地址 0xc00001c0b8

}

// 总结&分析
// defer 、 return 和 返回值 3 者的顺序是
// defer 多个defer是后进先出的原则，defer是 被调用函数执行返回前执行的。
// return 优先返回 ，defer 次之 ，最后返回值 返回给调用函数 的变量
// 值得说明的是 defer函数 的返回值 都是有参数名的。相应返回的变量赋值 例如 f1() return 0 => r = 0 , return r
// defer 函数的入参参数 与 外面的参数 是值传递 示例 f3 的 defer 外部 与 内部的 r 的内存地址 都是不同的
// 注意匿名函数 存在闭包 现象 例如 f1() 函数
