package ch108

func MainCall1() {
	demo11()
}

type Animal struct {
	Name string
	mean bool
}

type AnimalSounder interface {
	MakeNoise()
}

type Dog struct {
	Animal       //采用匿名变量的方式
	BarkStrength int
}

type Cat struct {
	Basics       Animal
	MeowStrength int
}

func demo11() {
	myDog := &Dog{
		Animal{
			"Rover", // Name
			false,   // mean
		},
		2, // BarkStrength
	}

	myCat := &Cat{
		Basics: Animal{
			Name: "Julius",
			mean: true,
		},
		MeowStrength: 3,
	}

	MakeSomeNoise(myDog)
	MakeSomeNoise(myCat)
}

// 进行了优化
//func (dog *Dog) MakeNoise() {
//	barkStrength := dog.BarkStrength
//
//	if dog.mean == true {
//		barkStrength = barkStrength * 5
//	}
//
//	for bark := 0; bark < barkStrength; bark++ {
//		fmt.Printf("BARK ")
//	}
//
//	fmt.Println("")
//}

//func (cat *Cat) MakeNoise() {
//	meowStrength := cat.MeowStrength
//
//	if cat.Basics.mean == true {
//		meowStrength = meowStrength * 5
//	}
//
//	for meow := 0; meow < meowStrength; meow++ {
//		fmt.Printf("MEOW ")
//	}
//
//	fmt.Println("")
//}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

// 总结&分析
// 知识点
// 1、接口 AnimalSounder 中 定义 MakeNoise ;接口不是可以实例化的类型，是行为声明，其他类型可以去实现接口声明的行为。
// 2、结构体 Dog、Cat ，其中 结构体的变量 采用了 变量声明和匿名 2种方式
// 3、结构体 Dog、Cat 都实现 AnimalSounder 接口的 MakeNoise 方法
// 4、继承体现 Dog、Cat 中 存在Animal的变量
// 5、在 Go 中，变量、结构体、成员、函数等的第一个字母的大小写决定了访问权限。大写字母开头表示公共的，可供外部调用；小写字母开头表示私有的，外部不能调用。
// 4、MakeSomeNoise 函数 是一个 AnimalSounder入参的函数，是通过接口实现多态体现
// 在Go中，通过方法实现接口的任何类型都表示接口类型，从而实现OOP的
