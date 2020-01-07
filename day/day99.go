package main

import "fmt"

//2 points
//1. the a awalys is 1
//2. a is new in for ,start with 1
func first_99() {
	a := 1
	for i := 0; i < 5; i++ {
		a := a + 1
		// a = a + 1 不一样！！！
		a *= 2
		fmt.Println(a)
	}
	fmt.Println(a)
}

//ret is shadowed during return  := 禁止新的变量on left
func second_99(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		// ret := 10
		ret = 10 //fixed
		return
	}
	return
}

func main() {
	//	first_99()
	result := second_99(10)
	fmt.Println(result)
}
