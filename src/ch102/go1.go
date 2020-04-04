package ch102

import (
	"fmt"
	"sync"
)

func MainCall1() {
	demo11()
}

var mu sync.Mutex
var chain string

func demo11() {
	chain = "main"
	A()
	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	//mu.Lock()  // A 函数已经加锁了，在锁没有释放前，C函数不能获取锁
	//defer mu.Unlock()
	chain = chain + " --> C"
}

// 总结&分析
// 锁的使用
// Mutex 锁，对共享资源进行管理，保证资源每次只能有一个访问/处理
