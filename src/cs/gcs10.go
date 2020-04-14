package cs

import "fmt"

func MainCall10() {
	main11()
}

type People1 interface {
	Speak1(string) string
}
type Std1 struct{}

func (stu *Std1) Speak1(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main11() {
	var peo People1 = &Std1{}
	think := "bitch"
	fmt.Println(peo.Speak1(think))
}
