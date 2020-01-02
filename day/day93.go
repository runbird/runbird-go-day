package main

import "fmt"

var x int

func init() {
	x++
}

//求两个数之间的较小值，能否在 该函数中添加一行代码将其功能补全。
func second_93(a int, b uint) {
	var min = 0
	fmt.Printf("the min of %d and %d is %d\n", a, b, min)
}

// copy(dst, src []Type) int
// The copy built-in function copies elements from a source slice into a
// destination slice. (As a special case, it also will copy bytes from a
// string to a slice of bytes.) The source and destination may overlap. Copy
// returns the number of elements copied, which will be the minimum of
// len(src) and len(dst).
func fix_93(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("the min of %d and %d is %d\n", a, b, min)
}

func main() {
	//init() init() 函数不能被其他函数调用，包括 main() 函数。
	second_93(1225, 256)
	fix_93(1225, 256)
	a := make([]struct{}, 15)
	b := copy([]int{1, 2, 3}, []int{2, 3, 3})
	fmt.Printf("{},{}", a, b)
}
