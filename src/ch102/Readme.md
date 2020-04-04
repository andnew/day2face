1.ch := make(chan interface{}) 和 ch := make(chan interface{},1) 有什么区别？
 答: 都是是 声明 一个 通道，类型是任意接口类型
    不同，第一个，容量是 0 ，第一个的容量是 1

2.下面的代码输出什么？请简要说明。 (D)
A. 不能编译；
B. 输出 main --> A --> B --> C；
C. 输出 main；
D. fatal error；

package main

import ch "ch102"

func main() {
	// 示例1
	ch.MainCall1()

	// defer Mutex 示例2
	//ch.MainCall2()

}
