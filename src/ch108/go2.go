package ch108

import "fmt"

func (animal *Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 5
	}

	for voice := 0; voice < strength; voice++ {
		fmt.Printf("%s ", sound)
	}

	fmt.Println("")
}

func (dog *Dog) MakeNoise() {
	dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (cat *Cat) MakeNoise() {
	cat.Basics.PerformNoise(cat.MeowStrength, "MEOW")
}

// 总结&分析
// 对 Dog、Cat 实现 AnimalSounder 的方法进行优化，将重复的代码 抽象到共用 PerformNoise 中
