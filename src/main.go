package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	//only del key once, and not del the current iteration key
	var o sync.Once
	for i := range m {
		o.Do(func() {
			for _, key := range []int{1, 2, 3} {
				if key != i {
					fmt.Printf("when iteration key %d, del key %d\n", i, key)
					delete(m, key)
					break
				}
			}
		})
		fmt.Printf("%d-->%d ", i, m[i])
	}

	fmt.Println("-------0x06 对 map 遍历时新增元素能遍历到么？--------")
	var createElemDuringIterMap = func() {
		var m = map[int]int{1: 1, 2: 2, 3: 3}
		for i := range m {
			m[4] = 4
			fmt.Printf("%d===%d ", i, m[i])
		}
	}
	for i := 0; i < 50; i++ {
		//some line will not show 44, some line will
		createElemDuringIterMap()
		fmt.Println()
	}
}

// 执行结果
