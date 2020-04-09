package main

import (
	"fmt"
	"os"
	"runtime"
)

//func main() {
//	runtime.GOMAXPROCS(1)
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//		}
//	}()
//
//	for {
//		fmt.Println("a")
//	}
//}

//可以通过阻塞的方式避免 CPU 占用，修复代码：
func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	select {}
}

// 总结&分析
// for {} 独占 CPU 资源导致其他 Goroutine 饿死。
