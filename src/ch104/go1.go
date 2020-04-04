package ch104

import (
	"sync"
	"time"
)

func MainCall1() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1) // 这些是多余的
	}()
	wg.Wait()
}

// 总结&分析
// 1、sync.WaitGroup 是一个计数器的功能，wg.Add 表示 计数器 增加1 ，wg.Done 计数器减值
