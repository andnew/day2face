package cs

import "fmt"

func MainCall1() {
	//case11()
	//case12()
	//case13()
	case14()
}

func case11() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")

}

func case12() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

//触发异常
//打印后
//打印中
//打印前

func case13() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {

		fmt.Println("打印后")
	}()

	panic("触发异常")
}

//执行结果
//打印后
//触发异常
//打印中
//打印前

func case14() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印前")
	}()

	defer func() { // 必须要先声明defer，否则recover()不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印中")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //err 就是panic传入的参数
		}
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

//执行结果
//触发异常
//打印后
//打印中
//打印前
