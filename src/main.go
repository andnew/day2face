package main

import "fmt"

type MyInt int

func (i MyInt) PrintInt() {
	fmt.Println(i)
}

func main1() {
	var i MyInt = 1
	i.PrintInt()
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
