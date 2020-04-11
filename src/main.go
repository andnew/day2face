package main

import (
	"encoding/json"
	"fmt"
)

//func main() {
//	var ch chan int
//	select {
//	case v, ok := <-ch:
//		println(v, ok)
//	default:
//		println("default")
//	}
//}

type People struct {
	Name string `json:"name"`
}

func main() {
	js := `{
         "name":"seekload"
    }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
