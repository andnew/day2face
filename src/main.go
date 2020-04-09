package main

import "fmt"

func main() {
	c := &ConfigOne{}
	c.String()
}

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

// 总结&分析
// 知识点：Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会被释放（内存逃逸），
// 因此在 Go 语言中返回函数参数或临时变量是安全的。
