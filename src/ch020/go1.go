package ch020

import "fmt"

func MainCall1() {
	call11()
}

type Person struct {
	age int
}

func call11() {
	person := &Person{28}
	fmt.Printf("--ch020--外--前---%p\n", &person)
	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Printf("---ch020--2---前--%p\n", &p)
		fmt.Println(p.age)
		fmt.Printf("---ch020--2---后--%p\n", &p)
	}(person)

	// 3.
	defer func() {
		fmt.Printf("---ch020--3---前--%p\n", &person)
		fmt.Println(person.age)
		fmt.Printf("---ch020--3---后--%p\n", &person)
	}()
	fmt.Printf("--ch020--外--后---%p\n", &person)
	person = &Person{29}
	fmt.Printf("--ch020--外--后2--%p\n", &person)

}

// 总结&分析
// 参考 ch018de的结论
