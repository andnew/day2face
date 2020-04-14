package main

import (
	"cs"
	"fmt"
	"runtime"
)

func main() {

	//1、写出下面代码输出内容
	//cs.MainCall1()
	//2、以下代码有什么问题，说明原因
	//cs.MainCall2()
	//3、下面的代码会输出什么，并说明原因
	//cs.MainCall3()
	//4、下面的代码会输出什么
	//cs.MainCall4()
	//5、下面代码会触发异常吗？请详细说明
	//cs.MainCall5()
	//6、下面代码输出什么？
	//cs.MainCall6()
	//7、请写出以下输入内容
	//cs.MainCall7()
	//8、下面的代码有什么问题
	//cs.MainCall8()
	//9、下面的迭代会有什么问题？
	//cs.MainCall9()
	//10、以下代码能编译过去吗？为什么？
	cs.MainCall10()
}

func init() {
	fmt.Println("Current Go Version:", runtime.Version())
}
