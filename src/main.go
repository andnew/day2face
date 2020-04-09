package main

func alwaysFalse() bool {
	return false
}
func main() {
	switch alwaysFalse(); {
	case true:
		println(true)
	case false:
		println(false)
	}
}

// 总结&分析
// 可以编译通过，输出：true。知识点：Go 代码断行规则。
