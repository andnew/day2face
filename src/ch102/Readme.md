1.ch := make(chan interface{}) 和 ch := make(chan interface{},1) 有什么区别？
 答: 都是是 声明 一个 通道，类型是任意接口类型
    不同，第一个，容量是 0 ，第一个的容量是 1
 官方,专业的回答 参考答案及解析：第一个是声明无缓存通道，第二个是声明缓存为 1 的通道。
   无缓存通道需要一直有接收者接收数据，写操作才会继续，不然会一直阻塞；
   而缓冲为 1 则即使没有接收者也不会阻塞，因为缓冲大小是 1 ，只有当放第二个值的时候，
   第一个还没被人拿走，这时候才会阻塞。注意这两者还是有区别的。

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
