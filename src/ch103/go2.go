package ch104

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

func MainCall2() {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
}

//下面代码输出什么？请简要说明。
// A. 不能编译；
//B. 输出 1；
//C. 程序 hang 住；
//D. fatal error；
// 参考答案及解析：D。当写锁阻塞时，新的读锁是无法申请的（有效防止写锁饥饿），导致死锁。
