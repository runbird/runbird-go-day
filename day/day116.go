package main

import "fmt"

var y int

func f116(x int) int {
	return 7
}

func first_116() {
	switch y = f116(2); { //switch case ;
	case y == 7: //case boolean :
		return
	}
}

//----------
func variadic116(ints ...int) []int {
	return ints
}

//可变函数。切片作为参数传入可变函数时不会创建新的切片。参见《可变函数》 https://mp.weixin.qq.com/s?__biz=MzI2MDA1MTcxMg==&mid=2648466706&idx=3&sn=003c213739e51088ad4947e473429775&chksm=f247437dc530ca6bafebe0a5a4090343cbf1eb992e36b6199cf213be6156273179465ed41348&scene=21#wechat_redirect
func second_116() {
	a := []int{1, 2, 3, 4}
	b := variadic116(a...)
	b[0], b[1] = b[1], b[0]
	fmt.Println(a)
}
func main() {
	second_116()
}
