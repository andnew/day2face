package ch106

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func MainCall1() {
	// 信道
	//demo11()
	// mutex
	//demo12()
	//
	demo13()
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
	print(a)
}

// 总结&分析
// 信道的特点

type MyMutex struct {
	count int
	sync.Mutex
}

func demo12() {
	var mu MyMutex
	var mu1 = mu
	mu.Lock()
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}

func demo13() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 10
	}()
	for i := 0; i < 10; i++ {
		go func(ch chan int) {
			time.Sleep(time.Second)
			<-ch
		}(ch)

	}

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
