package ch105

import (
	"fmt"
	"sync"
)

func MainCall2() {
	demo21()
}

type MyMutex struct {
	count int
	sync.Mutex
}

func demo21() {
	var mu MyMutex
	var mu1 = mu
	fmt.Println("mu的内存地址是  ", mu)
	fmt.Println("mu1的内存地址是 ", mu1)
	fmt.Println("mu == mu1 ", mu == mu1)
	mu.Lock()
	//var mu1 = mu // 锁内进行 变量赋值 导致 mu1 被锁，出现循环锁，正确的做法 ，这个代码向上提一行
	mu.count++
	fmt.Println("----mu.count++ 后 ", mu.count)
	fmt.Println("mu.count++后,mu == mu1 ", mu == mu1)
	mu.Unlock()
	mu1.Lock()
	fmt.Println("----mu1.count++ 前 ", mu1.count)
	mu1.count++
	fmt.Println("----mu1.count++ 后 ", mu1.count)
	fmt.Println("mu1.count++后,mu == mu1 ", mu == mu1)
	mu1.Unlock()
	fmt.Println(mu.count)
	fmt.Println(mu1.count)
}

// 总结&分析
// 1、Mutex 限制资源的共享， 通过加锁和解锁 控制的资源的并发访问
// 2、结构体的 是否执行同一个地址 ,比较通过  ==
