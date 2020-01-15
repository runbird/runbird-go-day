package main

import "fmt"

// 可变函数
// 可变参数函数，接收可变数量的参数的函数。
// 如果一个函数的最后一个参数的表示形如：…Type，则该参数（形参）可以接受不同数目的参数（实参）。

//func varib_1(elems ...Type) //注意，只允许最后一个参数是可变参数。
//func append(slice []Type, elems ...Type) []Type
//func append(slice []Type, elems ...Type) []Type
//s := []int{1,2,3}
//s = append(s,4,5,6,7,8)

//How to work?
func change(str1 string, s ...string) {
	fmt.Println(str1)
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
func executeChange() {
	slice := []string{"Hello", "World", "Go"}
	blog := "runbird"
	// change(blog, []string{slice}) //改变 change s ...string[]
	change(blog, slice...) //传参变为 slice...
}

func change2(s ...string) {
	s[0] = "runbird"
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
func executeChange2() {
	slice := []string{"Hello", "World", "Go"}
	change2(slice...)
	fmt.Println(slice)
}

func main() {
	executeChange()
}
