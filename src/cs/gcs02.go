package cs

import "fmt"

func MainCall2() {
	main21()
}

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, _ := range stus {
		stu := stus[i]
		fmt.Printf("%")
		m[stu.Name] = &stu
	}
	return m
}
func main21() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

// 总结&分析
// 利用for-range 进行数据赋值，for-range的中变量 是一个赋值的临时变脸，内存地址保持不变的
// 执行结果是
//key=zhou,value=&{wang 22}
//key=li,value=&{wang 22}
//key=wang,value=&{wang 22}
