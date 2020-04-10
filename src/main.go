package main

//func main() {
//
//loop:
//	fmt.Println("--------")
//
//	for i := 0; i < 10; i++ {
//	loop:
//		println(i)
//	}
//	goto loop
//}

// 总结&分析
// 编译报错 标签，不可以放在for内
// goto 不能跳转到其他函数或者内层代码。

//func main() {
//	/* 定义局部变量 */
//	var a int = 10
//
//	/* 循环 */
//LOOP:
//	for a < 20 {
//		if a == 15 {
//			/* 跳过迭代 */
//			a = a + 1
//			goto LOOP
//		}
//		fmt.Printf("a的值为 : %d\n", a)
//		a++
//	}
//}

// 总结&分析
// goto 与 标签组合 效果 continue 一致，跳出循环

func main() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2])
}

// 总结&分析
// 考查的知识点 有 defer 匿名函数 切片 指针
