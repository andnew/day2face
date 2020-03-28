package ch004

import "fmt"

func MainCall1() {

	list := make([]int, 5)
	list = append(list, 1)
	fmt.Println(list)
}

// 总结&分析
// 编译爆粗  first argument to append must be slice; have *[]int
// append 是 Go 的内置函数，append 作用在 切片 上
// 而 list 变量 是 数组
// 正确做法是 将 new 关键词 修改为 make
