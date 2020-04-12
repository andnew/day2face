package main

import "fmt"

func main1() {

	s1 := []int{1, 2, 3}
	s2 := s1[1:]             // {2, 3}
	s2[1] = 4                // {2, 4}
	fmt.Println(s1)          //{1,2,4}
	s2 = append(s2, 5, 6, 7) // {1,2,4,5,6,7}
	fmt.Println(s1)          // {1,2,4}
}

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
} // 结果是 1， 2
