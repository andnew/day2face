package main

import (
	"fmt"
	"sync"
	"time"
)

//type T struct{}
//
//func (*T) foo() {
//}
//func (T) bar() {
//}
//
//type S struct {
//	*T
//}
//
//func main() {
//	s := S{}
//	fmt.Printf("%#v\n",s) //main.S{T:(*main.T)(nil)}
//	_ = s.foo
//	s.foo()
//	_ = s.bar //s.bar 将被展开为 (*s.T).bar
//}
// 总结&分析
// panic: runtime error: invalid memory address or nil pointer dereference

//type data struct {
//	sync.Mutex
//}
//
//func (d data) test(s string) {// 修改为 func (d *data) test(s string) {
//	d.Lock()
//	defer d.Unlock()
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(s, i)
//		time.Sleep(time.Second)
//	}
//}
//
//func main() {
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//	var d data
//
//	go func() {
//		defer wg.Done()
//		d.test("read")
//	}()
//
//	go func() {
//		defer wg.Done()
//		d.test("write")
//	}()
//
//	wg.Wait()
//}
// 总结&分析
// 锁失效。将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。

type data struct {
	*sync.Mutex // *Mutex
}

func (d data) test(s string) { // 值方法
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	d := data{new(sync.Mutex)} // 初始化

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
