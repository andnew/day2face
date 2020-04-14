package cs

import "fmt"

func MainCall4() {
	main41()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main41() {
	t := Teacher{}
	//t.ShowA()
	t.ShowA()
	fmt.Println("---------------")
	t.People.ShowA()
}
