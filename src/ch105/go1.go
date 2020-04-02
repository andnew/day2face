package ch105

func MainCall1() {
	demo11()
}

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func demo11() {
	go f()
	c <- 0
	println("go1----", a)
}

// 总结&分析
// 1、协程
// 2、全局变量
