package main

import "fmt"

// 1.
func main1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[0 0 0 0 0 1 2 3]
}

// 2.
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s) //[1 2 3 4]
}

func funcMui(x, y int) (int, error) {
	return x + y, nil
}
