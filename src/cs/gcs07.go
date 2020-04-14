package cs

import "fmt"

func MainCall7() {
	main71()
}

func main71() {
	s := make([]int, 5)
	fmt.Printf("%p\n", s)
	s = append(s, 1, 2, 3)
	fmt.Printf("%p\n", s) //new pointer
	fmt.Println(s)
}
