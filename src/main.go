package main

import (
	"fmt"
	"sync"
)

func main() {

	// 方式一
	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		fmt.Printf("i:%d\n", i)
	//		wg.Done()
	//	}(i)
	//}

	//wg.Wait()

	//fmt.Println("exit")

	// 方式二
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	fmt.Println("exit")
}

// 总结&分析
// 知识点：Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放（内存逃逸），
// 因此在 Go 语言中返回函数参数或临时变量是安全的。
