package main

func main() {

	var s []int
	s = append(s, 1)

	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1
}

// 执行结果 panic: assignment to entry in nil map
