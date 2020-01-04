package main

import "fmt"

func main() {
	//Go提供了内置函数copy，可以讲一个切片复制到另一个切片。函数原型：
	//func copy(dst, src []Type) int
	//dst是目标切片，src是源切片，函数返回两者长度的最小值
	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := []int{4, 5, 6, 7}
	s4 := []int{1, 2, 3}
	// 1、
	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v, s2=%v\n", n1, s1, s2)
	fmt.Println("s1 == nil :", s1 == nil)
	// 2、
	n2 := copy(s2, s3)
	fmt.Printf("n2=%d, s2=%v, s3=%v\n", n2, s2, s3)
	// 3、
	n3 := copy(s3, s4)
	fmt.Printf("n3=%d, s3=%v, s4=%v\n", n3, s3, s4)
}
