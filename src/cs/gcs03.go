package cs

import (
	"fmt"
	"runtime"
	"sync"
)

func MainCall3() {
	main31()
}

func main31() {
	runtime.GOMAXPROCS(1)

	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count * 2)
	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("[%d]", i)
			wg.Done()
		}()
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Printf("-%d-", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
